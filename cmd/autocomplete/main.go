package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/gabivlj/candice/internals/autocomplete/server"
	"go.lsp.dev/jsonrpc2"
)

type rwc struct {
	r io.ReadCloser
	w io.WriteCloser
}

func (rwc *rwc) Read(b []byte) (int, error)  { return rwc.r.Read(b) }
func (rwc *rwc) Write(b []byte) (int, error) { return rwc.w.Write(b) }
func (rwc *rwc) Close() error {
	rwc.r.Close()
	return rwc.w.Close()
}

func connectLanguageServer(rwc io.ReadWriteCloser) jsonrpc2.Conn {
	bufStream := jsonrpc2.NewStream(rwc)
	rootConn := jsonrpc2.NewConn(bufStream)
	return rootConn
}

func main() {
	conn := connectLanguageServer(&rwc{os.Stdin, os.Stdout})
	ctx := context.TODO()

	server.InitializeHandlers()

	conn.Go(ctx, func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		m := req.Method()
		log.Println("method:", m, "-", "Handling...")
		data := req.Params()
		payload := map[string]interface{}{}
		rawJson, _ := data.MarshalJSON()
		json.Unmarshal(rawJson, &payload)
		log.Println(payload)
		if err := server.New(reply, req).HandleConnection(ctx); err != server.ErrHandlerForThisMethodDoesNotExist {
			return err
		} else if err == server.ErrHandlerForThisMethodDoesNotExist {
			log.Println("method", req.Method(), "doesn't have a handler")
		}

		return nil
	})
	<-conn.Done()
}
