package tree_printer

import (
	"strings"

	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
)

const FirstConnectorMultiple = "╦"
const StraightLine = "═"
const Wall = "║"
const MiddleConnector = "╠"
const LastConnector = "╚"

// const Connector

func Process(statement ast.Statement) string {
	switch t := statement.(type) {
	case *ast.ExpressionStatement:
		{

			return ConnectString("ast.ExpressionStatement", []string{processExpression(t.Expression)})
		}

	case *ast.DeclarationStatement:
		{
			return ConnectString("ast.DeclarationStatement", []string{ast.RetrieveID(t.Name), processExpression(t.Expression)})
		}

	case *ast.Block:
		{
			var statements []string
			for _, statement := range t.Statements {
				statements = append(statements, Process(statement))
			}
			return ConnectString("ast.Block", statements)
		}
	}

	return ""
}

func processExpression(expression ast.Expression) string {
	switch t := expression.(type) {
	case *ast.Identifier:
		{
			return t.String()
		}

	case *ast.Integer, *ast.Float:
		{
			return t.String() + " (" + t.GetType().String() + ")"
		}

	case *ast.AnonymousFunction:
		{
			label := "ast.AnonymousFunction"
			var types []string
			for i, ty := range t.GetFunctionType().Parameters {
				types = append(types, ast.RetrieveID(t.GetFunctionType().Names[i])+" "+ty.String())
			}
			tyParametersString := ConnectString("types.Parameters", types)
			s := Process(t.Block)
			returnType := t.GetFunctionType().Return
			if returnType == nil {
				returnType = ctypes.VoidType
			}

			tyReturnString := ConnectString("types.Return", []string{returnType.String()})
			return ConnectString(label, []string{tyParametersString, s, tyReturnString})
		}

	case *ast.BinaryOperation:
		{
			array := []string{
				processExpression(t.Left),
				t.Operation.String(),
				processExpression(t.Right),
			}
			return ConnectString("ast.BinaryOperation", array)
		}
	}
	return ""
}

func buildWall(space string, s string) string {
	wall := ""
	skipped := false
	for _, c := range strings.Split(s, "\n") {
		if !skipped {
			wall += c
			skipped = true
			continue
		}

		wall += Wall + c
	}

	if len(wall) == 0 {
		return wall
	}

	return wall[:len(wall)-1]
}

func getLine(pos int, len int, multipleBelow bool) string {
	if pos == 0 {
		if multipleBelow {
			return FirstConnectorMultiple
		}

		return StraightLine
	}

	if pos == len-1 {
		return LastConnector
	}

	return MiddleConnector
}

func getWall(multipleBelow bool) string {
	if multipleBelow {
		return Wall
	}

	return " "
}

func ConnectString(tag string, s []string) string {
	builder := strings.Builder{}
	spaces := strings.Repeat(" ", len(tag))
	for i, element := range s {
		toPrintFirst := spaces
		if i == 0 {
			toPrintFirst = tag
		}

		builder.WriteString(toPrintFirst + getLine(i, len(s), len(s) > 1))
		values := strings.Split(element, "\n")
		if len(values) > 0 {
			builder.WriteString(values[0] + "\n")
			wall := getWall(i != len(s)-1)
			for _, v := range values[1:] {
				builder.WriteString(spaces + wall + v + "\n")
			}
		}
	}

	return strings.TrimSpace(builder.String())
}
