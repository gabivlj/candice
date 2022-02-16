package semantic

import "github.com/gabivlj/candice/internals/ast"

func (s *Semantic) predefineTypes(statements []ast.Statement) {
	for _, statement := range statements {
		s.currentStatementBeingAnalyzed = statement
		switch t := statement.(type) {
		case *ast.StructStatement:
			{
				s.definedTypes[t.Type.Name] = t.Type
			}

		case *ast.MacroBlock:
			{
				s.predefineTypes(t.Statements)
			}
		}
	}
}

func (s *Semantic) predefineFunctions(statements []ast.Statement) {
	for _, statement := range statements {
		s.currentStatementBeingAnalyzed = statement
		switch t := statement.(type) {
		case *ast.FunctionDeclarationStatement:
			{
				s.replaceAnonymousFunctionParameterTypes(t.FunctionType)
				s.variables.Add(t.FunctionType.Name, s.newType(t.FunctionType))
			}

		case *ast.MacroBlock:
			{
				s.predefineFunctions(t.Statements)
			}
		}
	}
}

func (s *Semantic) fillTypes(statements []ast.Statement) {
	for _, statement := range statements {
		s.currentStatementBeingAnalyzed = statement
		switch t := statement.(type) {
		case *ast.TypeDefinition:
			{
				s.analyzeTypeDefinition(t)
			}

		case *ast.ExternStatement:
			{
				s.analyzeExternStatement(t)
			}

		case *ast.GenericTypeDefinition:
			{
				s.analyzeGenericTypeDefinition(t)
			}

		case *ast.ImportStatement:
			{
				s.analyzeImport(t)
			}
		case *ast.StructStatement:
			{
				s.analyzeStructStatement(t)
			}

		case *ast.MacroBlock:
			{
				s.fillTypes(t.Statements)
			}
		}
	}
}
