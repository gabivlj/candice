package server

import (
	"context"
	"errors"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
)

type Connection struct {
	reply jsonrpc2.Replier
	req   jsonrpc2.Request
}

var ErrHandlerForThisMethodDoesNotExist = errors.New("server: handler for this method doesn't exist")

var handlers map[string]func(context.Context, *Connection) error = map[string]func(context.Context, *Connection) error{}

func New(reply jsonrpc2.Replier, req jsonrpc2.Request) *Connection {
	return &Connection{
		reply,
		req,
	}
}

func (c *Connection) HandleConnection(ctx context.Context) error {
	handle, ok := handlers[c.req.Method()]
	if !ok {
		return ErrHandlerForThisMethodDoesNotExist
	}

	return handle(ctx, c)
}

func InitializeHandlers() {
	handlers[protocol.MethodTextDocumentDidChange] = func(_ context.Context, c *Connection) error {
		return c.onChange()
	}

	handlers[protocol.MethodInitialize] = func(ctx context.Context, c *Connection) error {
		return c.onInitialize(ctx)
	}
}
