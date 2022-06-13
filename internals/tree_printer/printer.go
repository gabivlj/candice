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

func ProcessProgram(program *ast.Program) string {
	var s []string
	for _, statement := range program.Statements {
		s = append(s, Process(statement))
	}

	return ConnectString("ast.Program", s)
}

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

	case *ast.BreakStatement:
		{
			return "ast.BreakStatement"
		}

	case *ast.ContinueStatement:
		{
			return "ast.ContinueStatement"
		}

	case *ast.AssignmentStatement:
		{
			return ConnectString("ast.AssignmentStatement", []string{processExpression(t.Left), processExpression(t.Expression)})
		}

	case *ast.ConditionPlusBlock:
		{
			return ConnectString("ast.ConditionPlusBlock", []string{processExpression(t.Condition), Process(t.Block)})
		}

	case *ast.ExternStatement:
		{
			return ConnectString("ast.ExternStatement", []string{t.Type.String()})
		}

	case *ast.ForStatement:
		{
			var strings []string
			if t.InitializerStatement != nil {
				strings = append(strings, "initializer: "+Process(t.InitializerStatement))
			}

			if t.Condition != nil {
				strings = append(strings, "condition: "+processExpression(t.Condition))
			}

			if t.Operation != nil {
				strings = append(strings, "operation: "+Process(t.Operation))
			}

			strings = append(strings, Process(t.Block))

			return ConnectString("ast.ForStatement", strings)
		}

	case *ast.IfStatement:
		{
			var strings []string
			strings = append(strings, processExpression(t.Condition))
			strings = append(strings, Process(t.Block))
			for _, elses := range t.ElseIfs {
				strings = append(strings, Process(elses))
			}

			if t.Else != nil {
				strings = append(strings, ConnectString("ast.Else", []string{Process(t.Else)}))
			}

			return ConnectString("ast.IfStatement", strings)
		}

	case *ast.FunctionDeclarationStatement:
		{
			var strings []string
			strings = append(strings, ast.RetrieveID(t.FunctionType.Name))
			parameterTypes, returnTypes := functionTypeToString(t.FunctionType)
			strings = append(strings, parameterTypes)
			strings = append(strings, returnTypes)
			strings = append(strings, Process(t.Block))
			return ConnectString("ast.FunctionDeclarationStatement", strings)
		}

	case *ast.GenericTypeDefinition:
		{
			return ConnectString("ast.GenericTypeDefinition", []string{ast.RetrieveID(t.Name)})
		}

	case *ast.ImportStatement:
		{
			var strings []string
			for _, t := range t.Types {
				strings = append(strings, t.String())
			}
			strings = append(strings, `"`+t.Path.String()+`"`)
			return ConnectString("ast.ImportStatement", strings)
		}

	case *ast.MacroBlock:
		{
			var s []string
			for _, statement := range t.Statements {
				s = append(s, Process(statement))
			}

			return ConnectString("ast.MacroBlock", s)
		}

	case *ast.ReturnStatement:
		{
			return ConnectString("ast.ReturnStatement", []string{processExpression(t.Expression)})
		}

	case *ast.StructStatement:
		{
			var strings []string
			for i, field := range t.Type.Fields {
				strings = append(strings, ConnectString("ast.Field", []string{ast.RetrieveID(t.Type.Names[i]), field.String()}))
			}

			return ConnectString("ast.StructStatement", strings)
		}

	case *ast.UnionStatement:
		{
			var strings []string
			for i, field := range t.Type.Fields {
				strings = append(strings, ConnectString("ast.Field", []string{ast.RetrieveID(t.Type.Names[i]), field.String()}))
			}

			return ConnectString("ast.UnionStatement", strings)
		}

	case *ast.SwitchStatement:
		{
			cases := []string{ConnectString("ast.Condition", []string{processExpression(t.Condition)})}
			for _, c := range t.Cases {
				cases = append(cases, ConnectString("ast.Case", []string{processExpression(c.Case), Process(c.Block)}))
			}

			if t.Default != nil {
				cases = append(cases, ConnectString("ast.Default", []string{Process(t.Default)}))
			}

			return ConnectString("ast.SwitchStatement", cases)
		}

	case *ast.TypeDefinition:
		{
			return ConnectString("ast.TypeDefinition", []string{ast.RetrieveID(t.Name), t.Type.String()})
		}

	case *ast.MultipleDeclarationStatement:
		{
			panic("unimplemented")
		}
	}

	return ""
}

func functionTypeToString(c *ctypes.Function) (string, string) {
	var types []string
	for i, ty := range c.Parameters {
		types = append(types, ast.RetrieveID(c.Names[i])+" "+ty.String())
	}

	tyParametersString := ConnectString("types.Parameters", types)
	returnType := c.Return
	if returnType == nil {
		returnType = ctypes.VoidType
	}

	tyReturnString := ConnectString("types.Return", []string{returnType.String()})
	return tyParametersString, tyReturnString
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
			tyParametersString, tyReturnString := functionTypeToString(t.FunctionType)
			s := Process(t.Block)
			return ConnectString(label, []string{tyParametersString, tyReturnString, s})
		}

	case *ast.BinaryOperation:
		{
			array := []string{
				processExpression(t.Left),
				"'" + t.Operation.String() + "'",
				processExpression(t.Right),
			}
			return ConnectString("ast.BinaryOperation", array)
		}

	case *ast.ArrayLiteral:
		{
			strings := []string{ConnectString("types.Array", []string{t.Type.String()})}
			for _, value := range t.Values {
				strings = append(strings, ConnectString("ast.ArrayElement", []string{processExpression(value)}))
			}

			return ConnectString("ast.ArrayLiteral", strings)
		}

	case *ast.BuiltinCall:
		{
			types := []string{}
			for _, t := range t.TypeParameters {
				types = append(types, t.String())
			}

			typesTree := ConnectString("ast.TypeParameters", types)
			parameters := []string{}
			for _, parameter := range t.Parameters {
				parameters = append(parameters, processExpression(parameter))
			}

			parametersTree := ConnectString("ast.Parameters", parameters)
			return ConnectString("ast.BuiltinCall", []string{ast.RetrieveID(t.Name), typesTree, parametersTree})
		}

	case *ast.Call:
		{
			parameters := []string{}
			for _, parameter := range t.Parameters {
				parameters = append(parameters, processExpression(parameter))
			}

			parametersTree := ConnectString("ast.Parameters", parameters)
			return ConnectString("ast.Call", []string{processExpression(t.Left), parametersTree})
		}
	}
	return ""
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
