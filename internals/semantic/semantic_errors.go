package semantic

import (
	"fmt"
	"strings"

	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/internals/token"
	"github.com/gabivlj/candice/pkg/logger"
)

type SemanticError struct {
	token   token.Token
	code    ast.Node
	message string
}

func (s *Semantic) checkWarningForMultipleStringAdding(binaryOperation *ast.BinaryOperation) {
	l, containsMoreToLeft := binaryOperation.Left.(*ast.BinaryOperation)
	r, containsMoreToRight := binaryOperation.Right.(*ast.BinaryOperation)

	if containsMoreToLeft || containsMoreToRight && (l != nil && l.Operation == ops.Add || r != nil && r.Operation == ops.Add) {
		s := binaryOperation.String()
		logger.Warning(fmt.Sprintf("You are adding more than 2 strings together, this can lead to a memory leak in your application because you can lose references to strings,"+
			" consider separating strings in different declarations.\n%s\n%s  Happening here", s, strings.Repeat("^", len(s))))
	}
}

func (s *Semantic) cantOperateThisOperationError(binaryOperation *ast.BinaryOperation, leftType ctypes.Type, rightType ctypes.Type) {
	line := fmt.Sprintf("%s", binaryOperation.String())
	markedLine := fmt.Sprintf("\t%s\n\t%s", line, strings.Repeat("^", len(line)))
	message := fmt.Sprintf("can't use a '%s' between a '%s' and a '%s'\n>>%s", binaryOperation.Operation, leftType, rightType, markedLine)
	s.error(message, binaryOperation.Token)
}
