package rpc

import (
	"encoding/json"
	"fmt"
)

var (
	vsn = "2.0"
)

type jsonrpcMessage struct {
	Error   *jsonError      `json:"error,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Version string          `json:"jsonrpc,omitempty"`
}

type jsonError struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

func (err *jsonError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("json-rpc error %d", err.Code)
	}
	return err.Message
}
