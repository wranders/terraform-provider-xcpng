package rpc

import "context"

type ServerCodec interface {
	readBatch() ([]*jsonrpcMessage, bool, error)
	close()
	jsonWriter
}

type jsonWriter interface {
	writeJSON(context.Context, interface{}) error
	closed() <-chan interface{}
	remoteAddr() string
}
