package node

import (
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/token"
)

type Node struct {
	Type ctypes.Type
	// TODO This will be filled by the lexer
	Token token.Token
}
