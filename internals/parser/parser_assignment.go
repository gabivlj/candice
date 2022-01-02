package parser

import (
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ops"
)

func (p *Parser) parsePossibleAssignment() ast.Statement {
	expr := p.parseExpression(0)
	if bin, ok := expr.(*ast.BinaryOperation); ok {
		if bin.Operation == ops.TempAssign {
			left, okLeft := bin.Left.(*ast.BinaryOperation)
			if okLeft {
				if left.Operation != ops.Dot {
					p.addErrorMessage("this operation is not permitted in an assignment")
				}
			} else {
				_, okAccess := bin.Left.(*ast.IndexAccess)
				if !okAccess {
					prefix, okPrefix := bin.Left.(*ast.PrefixOperation)
					if !okPrefix {
						_, okId := bin.Left.(*ast.Identifier)
						if !okId {
							p.addErrorMessage("this operation is not permitted in an assignment 2.0")
						}
					} else {
						if prefix.Operation != ops.Multiply {
							p.addErrorMessage("this prefix is not permitted in an assignment")
						}
						return &ast.AssignmentStatement{
							Left:       bin.Left,
							Expression: bin.Right,
						}
					}
				}
			}
			return &ast.AssignmentStatement{
				Left:       bin.Left,
				Expression: bin.Right,
			}
		}

		return &ast.ExpressionStatement{
			Token:      p.currentToken,
			Expression: expr,
		}
	}

	return &ast.ExpressionStatement{
		Token:      p.currentToken,
		Expression: expr,
	}
}
