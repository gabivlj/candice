package semantic

import (
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/token"
)

type SemanticError struct {
	token   token.Token
	code    ast.Node
	message string
}
