package server

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gabivlj/candice/internals/tree_printer"
	"go.lsp.dev/protocol"
)

func (c *Connection) onChange() error {
	js := c.req.Params()
	bytes, err := js.MarshalJSON()
	if err != nil {
		return err
	}

	var parameters protocol.DidChangeTextDocumentParams
	if err = json.Unmarshal(bytes, &parameters); err != nil {
		return err
	}

	txt := parameters.ContentChanges[0].Text
	log.Println(txt)
	log.Println("possible error writing output", tree_printer.WriteOutput(txt, os.Stderr))
	return nil
}
