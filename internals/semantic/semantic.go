package semantic

import (
	"errors"
	"fmt"
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/internals/token"
	"github.com/gabivlj/candice/internals/undomap"
	"github.com/gabivlj/candice/pkg/a"
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
	a.AssertEqual(key, "<main-frame>")
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
	case *ast.StructStatement:
		s.analyzeStructStatement(statementType)
		return
	}

	log.Fatalln("couldn't analyze statement: " + statement.String())
}

func (s *Semantic) analyzeStructStatement(statementType *ast.StructStatement) {
	s.definedTypes[statementType.Type.Name] = statementType.Type
	for _, t := range statementType.Type.Fields {
		unwrappedType := s.unwrap(t)
		if anonymous, ok := unwrappedType.(*ctypes.Anonymous); ok {
			definedType := s.unwrapAnonymous(anonymous)
			if definedType == statementType.Type && t == anonymous {
				s.error("recursive type detected", statementType.Token)
				return
			}
			s.swapTypes(t, definedType)
		}
	}
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
		s.variables.Add(declaration.Name, declType)
		return
	}

	// else we just prefer the one returned by analyzeExpression
	declaration.Type = ctype
	s.variables.Add(declaration.Name, ctype)
}

func (s *Semantic) unwrapAnonymous(t ctypes.Type) ctypes.Type {
	if anonymous, ok := t.(*ctypes.Anonymous); ok {
		return s.definedTypes[anonymous.Name]
	}

	return t
}

func (s *Semantic) unwrap(t ctypes.Type) ctypes.Type {
	if ptr, ok := t.(*ctypes.Pointer); ok {
		return s.unwrap(ptr.Inner)
	}

	if arr, ok := t.(*ctypes.Array); ok {
		return s.unwrap(arr.Inner)
	}

	return t
}

func (s *Semantic) swapTypes(t ctypes.Type, toSwap ctypes.Type) ctypes.Type {
	if ptr, ok := t.(*ctypes.Pointer); ok {
		val := s.swapTypes(ptr.Inner, toSwap)
		ptr.Inner = val
		return ptr
	}

	if arr, ok := t.(*ctypes.Array); ok {
		val := s.swapTypes(arr.Inner, toSwap)
		arr.Inner = val
		return arr
	}

	return toSwap
}

func (s *Semantic) areTypesEqual(first, second ctypes.Type) bool {
	if ctypes.IsPointer(first) && ctypes.IsPointer(second) {
		return s.areTypesEqual(first.(*ctypes.Pointer).Inner, second.(*ctypes.Pointer).Inner)
	}

	if ctypes.IsArray(first) {
		fArray := first.(*ctypes.Array)
		sArray, ok := second.(*ctypes.Array)
		if !ok {
			return false
		}

		return fArray.Length == sArray.Length && s.areTypesEqual(fArray.Inner, sArray.Inner)
	}

	return first == second
}

func (s *Semantic) analyzeExpression(expression ast.Expression) ctypes.Type {
	switch expressionType := expression.(type) {
	case *ast.Integer:
		return s.analyzeInteger(expressionType)
	case *ast.BuiltinCall:
		return s.analyzeBuiltinCall(expressionType)
	case *ast.BinaryOperation:
		return s.analyzeBinaryOperation(expressionType)
	case *ast.PrefixOperation:
		return s.analyzePrefixOperation(expressionType)
	case *ast.Identifier:
		return s.analyzeSimpleIdentifier(expressionType)
	default:
		log.Fatalln("couldn't analyze expression: " + expressionType.String())
	}
	return nil
}

func (s *Semantic) analyzeSimpleIdentifier(identifier *ast.Identifier) ctypes.Type {
	if identifierType := s.variables.Get(identifier.Name); identifierType != nil {
		return identifierType
	}
	s.error("undefined variable "+identifier.Name, identifier.Token)
	return ctypes.TODO()
}

func (s *Semantic) analyzePrefixOperation(prefixOperation *ast.PrefixOperation) ctypes.Type {
	t := s.analyzeExpression(prefixOperation.Right)
	if prefixOperation.Operation == ops.Bang || prefixOperation.Operation == ops.Add {
		if !ctypes.IsNumeric(t) {
			s.typeMismatchError(prefixOperation.String(), prefixOperation.Token, ctypes.LiteralToType("i32"), t)
		}

		return t
	}

	// We make this if because maybe in the future we don't want to '-' unsigned integers?
	if prefixOperation.Operation == ops.Subtract {
		if !ctypes.IsNumeric(t) {
			s.typeMismatchError(prefixOperation.String(), prefixOperation.Token, ctypes.LiteralToType("i32"), t)
		}

		return t
	}

	if prefixOperation.Operation == ops.BinaryAND {
		return &ctypes.Pointer{Inner: t}
	}

	if prefixOperation.Operation == ops.Multiply {
		if ptr, ok := t.(*ctypes.Pointer); !ok {
			s.typeMismatchError(prefixOperation.String(), prefixOperation.Token, &ctypes.Pointer{Inner: t}, t)
			return t
		} else {
			return ptr.Inner
		}
	}

	s.error("unknown prefix operator to analyze", prefixOperation.Token)

	return t
}

func (s *Semantic) analyzeBinaryOperation(binaryOperation *ast.BinaryOperation) ctypes.Type {
	op := binaryOperation.Operation
	if s.isArithmetic(op) {
		return s.analyzeArithmetic(binaryOperation)
	}
	s.error("can't analyze operator", binaryOperation.Token)
	return ctypes.TODO()
}

func (s *Semantic) isArithmetic(op ops.Operation) bool {
	return op == ops.OR || op == ops.Multiply || op == ops.BinaryXOR || op == ops.BinaryOR ||
		op == ops.BinaryAND || op == ops.AND || op == ops.Add || op == ops.Subtract || op == ops.LessThanEqual ||
		op == ops.LessThan || op == ops.Equals || op == ops.GreaterThan || op == ops.GreaterThanEqual ||
		op == ops.NotEquals || op == ops.Divide
}

func (s *Semantic) analyzeArithmetic(binaryOperation *ast.BinaryOperation) ctypes.Type {
	left := s.analyzeExpression(binaryOperation.Left)
	right := s.analyzeExpression(binaryOperation.Right)
	if !ctypes.IsNumeric(left) {
		s.error("expected numeric type, got: "+left.String(), binaryOperation.Token)
	}

	if !ctypes.IsNumeric(right) {
		s.error("expected numeric type, got: "+right.String(), binaryOperation.Token)
	}

	if !s.areTypesEqual(left, right) {
		s.typeMismatchError(binaryOperation.String(), binaryOperation.Token, right, left)
	}

	return left
}

func (s *Semantic) analyzeInteger(integer *ast.Integer) ctypes.Type {
	return integer.Type
}
