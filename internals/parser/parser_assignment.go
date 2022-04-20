package parser

import (
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/internals/token"
)

func (p *Parser) parsePossibleAssignment() ast.Statement {
	expr := p.parseExpression(0)
	// log.Println(expr)
	if bin, ok := expr.(*ast.BinaryOperation); ok {
		if bin.Operation == ops.TempAssign {
			switch left := bin.Left.(type) {
			case *ast.BinaryOperation:
				if left.Operation != ops.Dot {
					p.addErrorMessage("this operation is not permitted in an assignment")
				}
			case *ast.CommaExpressions:
				left.IsAssignment = true
			case *ast.IndexAccess:
			case *ast.PrefixOperation:
				if left.Operation != ops.Multiply {
					p.addErrorMessage("this prefix is not permitted in an assignment")
				}
			case *ast.Identifier:
			default:
				{
					p.addErrorMessage("this operation is not permitted in an assignment")
				}
			}

			return &ast.AssignmentStatement{
				Left:       bin.Left,
				Expression: bin.Right,
				Token:      p.currentToken,
			}
		}

		return &ast.ExpressionStatement{
			Token:      p.currentToken,
			Expression: expr,
		}
	}

	if commaExpr, ok := expr.(*ast.CommaExpressions); ok && p.currentToken.Type == token.COLON {
		declToken := p.nextToken()
		var t ctypes.Type = ctypes.TODO()
		if token.ASSIGN != p.currentToken.Type {
			t = p.parseType()
		}

		p.expect(token.ASSIGN)
		p.nextToken()
		return &ast.MultipleDeclarationStatement{
			Token:      declToken,
			Names:      ast.IdentifiersToStrings(commaExpr.Expressions),
			Type:       t,
			Expression: p.parseExpression(0),
			Constant:   false,
		}
	}

	return &ast.ExpressionStatement{
		Token:      p.currentToken,
		Expression: expr,
	}
}
