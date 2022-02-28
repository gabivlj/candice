package semantic

import (
	"fmt"
	"strings"

	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/token"
	"github.com/gabivlj/candice/pkg/logger"
)

type SemanticError struct {
	token   token.Token
	code    ast.Node
	message string
}

func (s *Semantic) checkWarningForMultipleStringAdding(binaryOperation *ast.BinaryOperation) {
	_, containsMoreToLeft := binaryOperation.Left.(*ast.BinaryOperation)
	_, containsMoreToRight := binaryOperation.Right.(*ast.BinaryOperation)
	if containsMoreToLeft || containsMoreToRight {
		s := binaryOperation.String()
		logger.Warning(fmt.Sprintf("You are adding more than 2 strings together, this can lead to a memory leak in your application because you can lose references to strings,"+
			" consider separating strings in different declarations.\n%s\n%s  Happening here", s, strings.Repeat("^", len(s))))
	}
}
