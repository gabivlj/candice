package compiler

import "github.com/gabivlj/candice/internals/ast"

// compileStructTypes tries to compile only type definitions and structs
func (c *Compiler) compileStructTypes(statements []ast.Statement) {
	for _, statement := range statements {
		switch t := statement.(type) {
		case *ast.TypeDefinition:
			c.compileTypeDefinition(t)

		case *ast.StructStatement:
			c.compileStruct(t)

		case *ast.UnionStatement:
			c.compileUnion(t)

		case *ast.MacroBlock:
			c.compileStructTypes(t.Statements)
		}
	}
}

// compileFunctionTypes makes a compile pass that defines function types without
// compiling their body
func (c *Compiler) compileFunctionTypes(statements []ast.Statement) {
	for _, statement := range statements {
		switch t := statement.(type) {
		case *ast.FunctionDeclarationStatement:
			c.compileFunctionType(t.FunctionType.Name, t)

		case *ast.MacroBlock:
			c.compileFunctionTypes(t.Statements)
		}
	}
}
