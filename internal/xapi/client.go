package xapi

import (
	"net/http"
	"net/url"

	"github.com/wranders/terraform-provider-xcpng/internal/rpc"
)

type Client struct {
	rpc *rpc.Client

	Session Session
	VM      VM
}

func NewClient(rawurl string, transport *http.Transport) (*Client, error) {
	var u *url.URL
	var err error
	var repErr error
	u, err = url.ParseRequestURI(rawurl)
	if err != nil || u.Host == "" {
		u, repErr = url.ParseRequestURI("http://" + rawurl)
		if repErr != nil {
			return nil, err
		}
	}
	u.Path = "jsonrpc"

	if transport == nil {
		transport = &http.Transport{}
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	rpc, err := rpc.DialHTTPWithClient(u.String(), httpClient)
	if err != nil {
		return nil, err
	}

	return prepareClient(rpc), nil
}

func prepareClient(rpc *rpc.Client) *Client {
	var client Client
	client.rpc = rpc

	client.Session = Session{&client}
	client.VM = VM{&client}

	return &client
}
