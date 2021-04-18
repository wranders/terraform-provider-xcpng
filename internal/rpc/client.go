package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"sync/atomic"
)

var (
	ErrClientQuit = errors.New("client is closed")
	ErrNoResult   = errors.New("no result in JSON-RPC response")
)

type Client struct {
	idCounter uint32
	idgen     func() ID
	isHTTP    bool
	writeConn jsonWriter
}

type reconnectFunc func(context.Context) (ServerCodec, error)

type requestOp struct {
	ids  []json.RawMessage
	err  error
	resp chan *jsonrpcMessage
}

func (op *requestOp) wait(ctx context.Context, c *Client) (*jsonrpcMessage, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case resp := <-op.resp:
		return resp, op.err
	}
}

func newClient(initctx context.Context, connect reconnectFunc) (*Client, error) {
	conn, err := connect(initctx)
	if err != nil {
		return nil, err
	}
	c := initClient(conn, randomIDGenerator())
	return c, nil
}

func initClient(conn ServerCodec, idgen func() ID) *Client {
	_, isHTTP := conn.(*httpConn)
	c := &Client{
		idgen:     idgen,
		isHTTP:    isHTTP,
		writeConn: conn,
	}
	return c
}

func (c *Client) nextID() json.RawMessage {
	id := atomic.AddUint32(&c.idCounter, 1)
	return strconv.AppendUint(nil, uint64(id), 10)
}

func (c *Client) Call(result interface{}, method string, args ...interface{}) error {
	ctx := context.Background()
	return c.CallContext(ctx, result, method, args...)
}

func (c *Client) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	if result != nil && reflect.TypeOf(result).Kind() != reflect.Ptr {
		return fmt.Errorf("call result parameter must be pointer or nil interface: %v", result)
	}
	msg, err := c.newMessage(method, args...)
	if err != nil {
		return err
	}
	op := &requestOp{ids: []json.RawMessage{msg.ID}, resp: make(chan *jsonrpcMessage, 1)}

	err = c.sendHTTP(ctx, op, msg)
	if err != nil {
		return err
	}

	switch resp, err := op.wait(ctx, c); {
	case err != nil:
		return err
	case resp.Error != nil:
		return resp.Error
	case len(resp.Result) == 0:
		return ErrNoResult
	default:
		return json.Unmarshal(resp.Result, &result)
	}
}

func (c *Client) newMessage(method string, paramsIn ...interface{}) (*jsonrpcMessage, error) {
	msg := &jsonrpcMessage{Version: vsn, ID: c.nextID(), Method: method}
	if paramsIn != nil {
		var err error
		if msg.Params, err = json.Marshal(paramsIn); err != nil {
			return nil, err
		}
	}
	return msg, nil
}
