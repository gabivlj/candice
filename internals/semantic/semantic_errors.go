package semantic

import (
	"fmt"
	"strings"

	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/internals/token"
	"github.com/gabivlj/candice/pkg/format"
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

func (s *Semantic) checkDereferenceOnArithmeticErrors(left, right ctypes.Type, binary *ast.BinaryOperation, prioritiseDereferenceFix bool) {
	leftExpr := binary.Left
	rightExpr := binary.Right
	leftUnwrapped, depthLeft := ctypes.UnwrapPossiblePointerAndDepth(left)
	rightUnwrapped, depthRight := ctypes.UnwrapPossiblePointerAndDepth(right)
	if depthLeft == 0 && depthRight == 0 {
		return
	}

	if !s.areTypesEqual(leftUnwrapped, rightUnwrapped) {
		if ctypes.IsNumeric(leftUnwrapped) && ctypes.IsNumeric(rightUnwrapped) {

		} else {
			return
		}
	}

	if ctypes.IsArray(left) || ctypes.IsArray(right) {
		return
	}

	message := fmt.Sprintf("%s differ on their pointer depth, (%s ≠ %s)\n", binary, left, right)

	diff := depthRight - depthLeft
	if diff < 0 {
		diff = -diff
	}

	message += s.getCurrentStatementLineFormatted()
	if depthLeft > depthRight {
		if prioritiseDereferenceFix {
			newLeft := s.wrapInOperators(leftExpr, diff, ops.Multiply)
			binary.Left = newLeft
			message += fmt.Sprintf("\nHint: consider dereferencing '%s' as '%s'\n", leftExpr, newLeft)
		} else {
			newRight := s.wrapInOperators(rightExpr, diff, ops.Reference)
			binary.Right = newRight
			message += fmt.Sprintf("\nHint: consider referencing '%s' as '%s'\n", rightExpr, newRight)
		}
	} else {
		if !prioritiseDereferenceFix {
			newLeft := s.wrapInOperators(leftExpr, diff, ops.Reference)
			binary.Left = newLeft
			message += fmt.Sprintf("\nHint: consider referencing '%s' as '%s'\n", leftExpr, newLeft)
		} else {
			newRight := s.wrapInOperators(rightExpr, diff, ops.Multiply)
			binary.Right = newRight
			message += fmt.Sprintf("\nHint: consider dereferencing '%s' as '%s'\n", rightExpr, newRight)
		}
	}
	message += s.getCurrentStatementLineFormatted()
	s.error(message, leftExpr.GetToken())
}

func (s *Semantic) checkDereferenceErrors(line string, left, right ctypes.Type, leftExpr ast.Expression) {
	leftUnwrapped, depthLeft := ctypes.UnwrapPossiblePointerAndDepth(left)
	rightUnwrapped, depthRight := ctypes.UnwrapPossiblePointerAndDepth(right)
	if !s.areTypesEqual(leftUnwrapped, rightUnwrapped) {
		return
	}

	if ctypes.IsArray(left) || ctypes.IsArray(right) {
		return
	}

	message := fmt.Sprintf("%s differs on the expected type pointer depth, (%s ≠ %s)\n", leftExpr, left, right)

	diff := depthRight - depthLeft
	if diff < 0 {
		diff = -diff
	}

	message += s.getCurrentStatementLineFormatted()
	if depthLeft > depthRight {
		newLeft := s.wrapInOperators(leftExpr, diff, ops.Multiply)
		message += fmt.Sprintf("\nHint: consider dereferencing '%s' as '%s'\n", leftExpr, newLeft)
	} else {
		newRight := s.wrapInOperators(leftExpr, diff, ops.Reference)
		message += fmt.Sprintf("\nHint: consider referencing '%s' as '%s'\n", leftExpr, newRight)
	}

	s.error(message, leftExpr.GetToken())
}

func (semantic *Semantic) getCurrentStatementLineFormatted() string {
	if semantic.currentStatementBeingAnalyzed != nil {
		s := semantic.currentStatementBeingAnalyzed.String()
		return semantic.formatLine(s)
	}

	return ""
}

func (semantic *Semantic) formatLine(s string) string {
	elements := strings.Split(format.StringWithTabs(s, 1), "\n")
	currentStatementLen := len(elements[0])
	return fmt.Sprintf("\n\t>> %s\n\t   %s\n%s", elements[0], strings.Repeat("^", currentStatementLen), strings.Join(elements[1:], "\n"))
}

func (s *Semantic) typeMismatchError(node string, wrongPart ast.Expression, tok token.Token, expected, got ctypes.Type) {
	var message string
	s.checkDereferenceErrors(node, got, expected, wrongPart)

	if len(s.Errors) == 0 {
		message = fmt.Sprintf("\n\n%s \n%s mismatched types, expected %s, got %s\n", node, strings.Repeat("^", len(node)), expected.String(), got.String())
	} else {
		message = fmt.Sprintf("can't recover from the errors\n")
	}

	if wrongPart != nil && len(s.Errors) == 0 {
		left := wrongPart.String()
		message += fmt.Sprintf("Hint: maybe are you missing a cast here?\n%s\n%s\n", left, strings.Repeat("^", len(left)))
	} else if len(s.Errors) > 0 {
		message += fmt.Sprintf("Hint: check compiler errors above.")
	}

	s.error(message, tok)
}

// Type mismatch for an arithmetic operation like '+' excepting '<<' and '>>',
// this method can do reference and pointer hints.
func (s *Semantic) typeMismatchBlameArithmeticExpressionError(node string, binary *ast.BinaryOperation, tok token.Token, blameRight bool) {
	s.checkDereferenceOnArithmeticErrors(binary.Left.GetType(), binary.Right.GetType(), binary, true)
	expected := binary.Left.GetType()
	got := binary.Right.GetType()
	if s.areTypesEqual(expected, got) {
		return
	}

	s.typeMismatchBlameBinaryExpressionError(binary.String(), binary, tok, expected, got, blameRight)
}

// Type mismatch for every binary operation
func (s *Semantic) typeMismatchBlameBinaryExpressionError(node string, binary *ast.BinaryOperation, tok token.Token, expected ctypes.Type, got ctypes.Type, blameRight bool) {
	s.checkDereferenceOnArithmeticErrors(binary.Left.GetType(), binary.Right.GetType(), binary, true)
	wrongLine := s.getCurrentStatementLineFormatted()
	// Decide which side of the operation do you want to cast
	if blameRight {
		binary.Right = &ast.BuiltinCall{
			Name:           "cast",
			TypeParameters: []ctypes.Type{expected},
			Parameters:     []ast.Expression{binary.Right},
		}
	} else {
		binary.Left = &ast.BuiltinCall{
			Name:           "cast",
			TypeParameters: []ctypes.Type{got},
			Parameters:     []ast.Expression{binary.Left},
		}
	}
	message := fmt.Sprintf("\n%s\n\n%s \n%s mismatched types, expected %s, got %s\n", wrongLine, node, strings.Repeat("^", len(node)), expected.String(), got.String())
	if ctypes.IsNumeric(expected) && ctypes.IsNumeric(got) {
		message += fmt.Sprintf("Try doing:\n- [%d:%d] %s\n", binary.GetToken().Line, binary.GetToken().Position, binary.String())
	}

	s.error(message, tok)
}

func (s *Semantic) throwInvalidOperationForConstant(message string, node ast.Node) {
	if !s.expectConstantExpression {
		return
	}

	s.error(fmt.Sprintf("%sthis is an invalid operation for a constant expression\n\t%s", message, s.formatLine(node.String())), node.GetToken())
}
