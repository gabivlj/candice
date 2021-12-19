package semantic

import (
	"errors"
	"fmt"
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/token"
	"github.com/gabivlj/candice/internals/undomap"
	"log"
)

type Semantic struct {
	variables       *undomap.UndoMap
	definedTypes    map[string]ctypes.Type
	builtinHandlers map[string]func(builtin *ast.BuiltinCall) ctypes.Type
	errors          []error
}

func New() *Semantic {
	s := &Semantic{
		variables:       undomap.New(),
		definedTypes:    map[string]ctypes.Type{},
		builtinHandlers: map[string]func(builtin *ast.BuiltinCall) ctypes.Type{},
		errors:          []error{},
	}
	s.builtinHandlers["cast"] = s.analyzeCast
	return s
}

func (s *Semantic) enterFrame() {
	s.variables.Add("<main-frame>", ctypes.TODO())
}

func (s *Semantic) leaveFrame() {
	key := ""
	for key != "<main-frame>" {
		key, _ = s.variables.Pop()
	}
}

func (s *Semantic) error(msg string, tok token.Token) {
	s.errors = append(s.errors, errors.New(fmt.Sprintf("error analyzing on %d:%d (at %s): %s", tok.Line, tok.Position, tok.Type, msg)))
}

func (s *Semantic) typeMismatchError(node string, tok token.Token, expected, got ctypes.Type) {
	message := fmt.Sprintf("_%s_ :: mismatched types, expected=%s, got=%s", node, expected.String(), got.String())
	s.error(message, tok)
}

func (s *Semantic) Analyze(program *ast.Program) {
	s.enterFrame()
	for _, statement := range program.Statements {
		s.analyzeStatement(statement)
		if len(s.errors) > 0 {
			return
		}
	}
	s.leaveFrame()
}

func (s *Semantic) analyzeStatement(statement ast.Statement) {
	switch statementType := statement.(type) {
	case *ast.DeclarationStatement:
		s.analyzeDeclarationStatement(statementType)
		return
	}

	log.Fatalln("couldn't analyze statement: " + statement.String())
}

func (s *Semantic) analyzeBuiltinCall(call *ast.BuiltinCall) ctypes.Type {
	if builtinHandler, ok := s.builtinHandlers[call.Name]; ok {
		return builtinHandler(call)
	}
	s.error("unknown builtin call", call.Token)
	return ctypes.TODO()
}

func (s *Semantic) analyzeDeclarationStatement(declaration *ast.DeclarationStatement) {
	ctype := s.analyzeExpression(declaration.Expression)
	declType := declaration.Type

	// Check if declaration is forcing the type
	if declType != ctypes.TODO() {
		if !s.areTypesEqual(declType, ctype) {
			s.typeMismatchError(declaration.String(), declaration.Token, declType, ctype)
			return
		}

		return
	}

	// else we just prefer the one returned by analyzeExpression
	declaration.Type = ctype

}

func (s *Semantic) areTypesEqual(first, second ctypes.Type) bool {
	return first == second
}

func (s *Semantic) analyzeExpression(expression ast.Expression) ctypes.Type {
	switch expressionType := expression.(type) {
	case *ast.Integer:
		return s.analyzeInteger(expressionType)
	case *ast.BuiltinCall:
		return s.analyzeBuiltinCall(expressionType)
	default:
		log.Fatalln("couldn't analyze expression: " + expressionType.String())
	}
	return nil
}

func (s *Semantic) analyzeInteger(integer *ast.Integer) ctypes.Type {
	return integer.Type
}
