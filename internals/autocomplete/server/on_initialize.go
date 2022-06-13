package server

import (
	"context"

	"go.lsp.dev/protocol"
)

func (c *Connection) onInitialize(ctx context.Context) error {
	return c.reply(ctx, protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			CompletionProvider: &protocol.CompletionOptions{
				TriggerCharacters: []string{"."},
			},
			DefinitionProvider: true,
			HoverProvider:      true,
			// SignatureHelpProvider:
			TextDocumentSync: protocol.TextDocumentSyncOptions{
				Change:    protocol.TextDocumentSyncKindFull,
				OpenClose: true,
				Save: &protocol.SaveOptions{
					IncludeText: true,
				},
			},
		},
	}, nil)
}
