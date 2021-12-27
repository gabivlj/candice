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
	"strconv"
)

type Semantic struct {
	variables                 *undomap.UndoMap
	definedTypes              map[string]ctypes.Type
	builtinHandlers           map[string]func(builtin *ast.BuiltinCall) ctypes.Type
	returns                   bool
	currentExpectedReturnType ctypes.Type
	Errors                    []error
}

func New() *Semantic {
	s := &Semantic{
		variables:                 undomap.New(),
		definedTypes:              map[string]ctypes.Type{},
		builtinHandlers:           map[string]func(builtin *ast.BuiltinCall) ctypes.Type{},
		Errors:                    []error{},
		currentExpectedReturnType: ctypes.VoidType,
		returns:                   false,
	}

	s.builtinHandlers["cast"] = s.analyzeCast
	s.builtinHandlers["alloc"] = s.analyzeAlloc
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
	s.Errors = append(s.Errors, errors.New(fmt.Sprintf("error analyzing on %d:%d (at %s): %s", tok.Line, tok.Position, tok.Type, msg)))
}

func (s *Semantic) typeMismatchError(node string, tok token.Token, expected, got ctypes.Type) {
	message := fmt.Sprintf("_%s_ :: mismatched types, expected=%s, got=%s", node, expected.String(), got.String())
	s.error(message, tok)
}

func (s *Semantic) Analyze(program *ast.Program) {
	s.enterFrame()
	for _, statement := range program.Statements {
		s.analyzeStatement(statement)
		if len(s.Errors) > 0 {
			return
		}
	}
	s.leaveFrame()
}

func (s *Semantic) analyzeStatement(statement ast.Statement) {

	if statement == nil {
		return
	}

	switch statementType := statement.(type) {
	case *ast.DeclarationStatement:
		s.analyzeDeclarationStatement(statementType)
		return
	case *ast.StructStatement:
		s.analyzeStructStatement(statementType)
		return
	case *ast.BreakStatement:
		return
	case *ast.IfStatement:
		s.analyzeIfStatement(statementType)
		return
	case *ast.ForStatement:
		s.analyzeForStatement(statementType)
		return
	case *ast.FunctionDeclarationStatement:
		s.analyzeFunctionStatement(statementType)
		return
	case *ast.AssignmentStatement:
		s.analyzeAssigmentStatement(statementType)
		return

	case *ast.ExpressionStatement:
		s.analyzeExpression(statementType.Expression)
		return

	case *ast.ReturnStatement:
		// NOTE the idea here would be having a "current" expected return type,
		// when we find a return we check with the expected return type.
		// If they don't match we add an error.
		// When analyzing a block we would analyze if all its
		// branches return true. Finding a return type guarantees that on a block it will return something.
		// Then when someone analyzes that block will guess that it always returns something.
		// When we find an if, we need to check that every block analyzed returns something (an else must exist)
		// When we find a block we analyze statement and check if the return flag is true, then statement means that this
		// block always returns.
		// ```
		// for analyzeStmt(), if thisStatementReturns: return...
		// if EOF thisStatementReturns = false
		//
		// ```
		s.analyzeReturnStatement(statementType)
		return
	}

	log.Fatalln("couldn't analyze statement: " + statement.String() + " ")
}

func (s *Semantic) analyzeAssigmentStatement(assign *ast.AssignmentStatement) {
	right := s.analyzeExpression(assign.Expression)
	left := s.analyzeExpression(assign.Left)
	if !s.areTypesEqual(left, right) {
		//todo: token
		s.typeMismatchError(assign.String(), token.Token{}, left, right)
	}
}

func (s *Semantic) analyzeBlock(block *ast.Block) {
	s.enterFrame()
	for _, stmt := range block.Statements {
		if s.returns {
			return
		}
		s.analyzeStatement(stmt)
	}
	s.leaveFrame()
}

func (s *Semantic) analyzeForStatement(forStatement *ast.ForStatement) {
	s.enterFrame()

	s.analyzeStatement(forStatement.InitializerStatement)

	condition := s.analyzeExpression(forStatement.Condition)

	if !ctypes.IsNumeric(condition) && condition != ctypes.VoidType {
		s.typeMismatchError(forStatement.Condition.String(), forStatement.Token, ctypes.I32, condition)
	}

	s.analyzeStatement(forStatement.Operation)

	s.analyzeBlock(forStatement.Block)

	s.leaveFrame()
}

func (s *Semantic) analyzeFunctionStatement(fun *ast.FunctionDeclarationStatement) {
	if fun.FunctionType.Return == nil {
		fun.FunctionType.Return = ctypes.VoidType
	}

	s.variables.Add(fun.FunctionType.Name, fun.FunctionType)

	s.enterFrame()

	for i, param := range fun.FunctionType.Parameters {
		s.variables.Add(fun.FunctionType.Names[i], param)
	}

	temporaryExpectedReturnType := s.currentExpectedReturnType
	s.currentExpectedReturnType = fun.FunctionType.Return

	if fun.Block != nil {
		for _, statement := range fun.Block.Statements {
			if s.returns {
				break
			}
			s.analyzeStatement(statement)
		}
	}

	if !s.returns && fun.FunctionType.Return != ctypes.VoidType {
		s.error("not all paths of the function '"+fun.FunctionType.String()+"'  return a variable", fun.Token)
	}

	s.returns = false
	s.leaveFrame()

	s.currentExpectedReturnType = temporaryExpectedReturnType
}

func (s *Semantic) analyzeIfStatement(ifStatement *ast.IfStatement) {
	condition := s.analyzeExpression(ifStatement.Condition)
	if !ctypes.IsNumeric(condition) {
		s.typeMismatchError(ifStatement.Condition.String(), ifStatement.Token, ctypes.I32, condition)
	}

	doesReturn := true
	s.analyzeBlock(ifStatement.Block)
	if !s.returns {
		doesReturn = false
		s.returns = false
	}

	for _, currentIf := range ifStatement.ElseIfs {
		condition := s.analyzeExpression(currentIf.Condition)
		if !ctypes.IsNumeric(condition) {
			//todo
			s.typeMismatchError(currentIf.Condition.String(), token.Token{}, ctypes.I32, condition)
		}
		s.analyzeBlock(currentIf.Block)
		if !s.returns {
			doesReturn = false
			s.returns = false
		}
	}

	if ifStatement.Else != nil {
		s.analyzeBlock(ifStatement.Else)
		if !s.returns {
			doesReturn = false
			s.returns = false
		}
	} else {
		doesReturn = false
	}

	s.returns = doesReturn
}

func (s *Semantic) analyzeReturnStatement(returnStatement *ast.ReturnStatement) {
	theType := s.analyzeExpression(returnStatement.Expression)
	if !s.areTypesEqual(theType, s.currentExpectedReturnType) {
		s.typeMismatchError(returnStatement.String(), returnStatement.Token, s.currentExpectedReturnType, theType)
	}
	s.returns = true
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

	if toSwap == ctypes.TODO() {
		trueType := s.unwrapAnonymous(t)
		if _, ok := trueType.(*ctypes.Anonymous); (ok && trueType == t) || trueType == nil {
			s.error("unknown type "+t.String(), token.Token{})
		}
		return trueType
	}

	return toSwap
}

func (s *Semantic) areTypesEqual(first, second ctypes.Type) bool {
	first = s.unwrapAnonymous(first)
	second = s.unwrapAnonymous(second)

	if ctypes.IsPointer(first) && ctypes.IsPointer(second) {
		return s.areTypesEqual(first.(*ctypes.Pointer).Inner, second.(*ctypes.Pointer).Inner)
	}

	if firstFunc, ok := first.(*ctypes.Function); ok {
		secondFunc, ok := second.(*ctypes.Function)
		if !ok {
			return false
		}
		for i, param := range firstFunc.Parameters {
			if !s.areTypesEqual(secondFunc.Parameters[i], param) {
				return false
			}
		}
		return s.areTypesEqual(firstFunc.Return, secondFunc.Return)
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
	if expression == nil {
		return ctypes.VoidType
	}

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
	case *ast.Call:
		return s.analyzeFunctionCall(expressionType)
	case *ast.ArrayLiteral:
		return s.analyzeArrayLiteral(expressionType)
	case *ast.StructLiteral:
		return s.analyzeStructLiteral(expressionType)
	case *ast.IndexAccess:
		return s.analyzeIndexAccess(expressionType)
	default:
		log.Fatalln("couldn't analyze expression: " + expressionType.String())
	}
	return nil
}

func (s *Semantic) analyzeIndexAccess(indexAccess *ast.IndexAccess) ctypes.Type {
	leftType := s.analyzeExpression(indexAccess.Left)

	if !ctypes.IsArray(leftType) && !ctypes.IsPointer(leftType) {
		s.error("expected a pointer or an array for an index access, instead we got "+leftType.String(), indexAccess.Token)
	}

	indexType := s.analyzeExpression(indexAccess.Access)
	if !ctypes.IsNumeric(indexType) {
		s.typeMismatchError(indexAccess.String(), indexAccess.Token, ctypes.I32, indexType)
	}

	if arr, ok := leftType.(*ctypes.Array); ok {
		return arr.Inner
	}

	if ptr, ok := leftType.(*ctypes.Pointer); ok {
		return ptr.Inner
	}

	s.error("mismatched types on index access, internal compiler bug", indexAccess.Token)

	return ctypes.TODO()
}

func (s *Semantic) analyzeStructLiteral(structLiteral *ast.StructLiteral) ctypes.Type {
	possibleStructType, ok := s.definedTypes[structLiteral.Name]

	if !ok {
		s.error("undefined struct "+structLiteral.Name+": "+structLiteral.String(), structLiteral.Token)
		return ctypes.TODO()
	}

	structType, ok := s.unwrapAnonymous(possibleStructType).(*ctypes.Struct)

	if !ok {
		s.error("undefined struct "+structLiteral.Name+": "+structLiteral.String(), structLiteral.Token)
		return ctypes.TODO()
	}

	structLiteral.Type = structType

	paramMap := map[string]int{}
	for i, name := range structType.Names {
		paramMap[name] = i
	}

	for _, value := range structLiteral.Values {
		index, ok := paramMap[value.Name]
		if !ok {
			s.error("undefined attribute on struct literal "+value.Name, structLiteral.Token)
		}
		expression := s.analyzeExpression(value.Expression)
		if !s.areTypesEqual(structType.Fields[index], expression) {
			s.typeMismatchError(structLiteral.String(), structLiteral.Token, structType.Fields[index], expression)
		}
	}

	return structType
}

func (s *Semantic) analyzeArrayLiteral(arrayLiteral *ast.ArrayLiteral) ctypes.Type {
	arrayType := arrayLiteral.Type.(*ctypes.Array)
	currType := arrayType.Inner

	if int(arrayType.Length) < len(arrayLiteral.Values) {
		s.error("expected an array of length "+strconv.FormatInt(arrayType.Length, 10)+" or less", arrayLiteral.Token)
	}

	for _, expr := range arrayLiteral.Values {
		t := s.analyzeExpression(expr)
		if !s.areTypesEqual(currType, t) {
			s.typeMismatchError(arrayLiteral.String(), arrayLiteral.Token, currType, t)
		}
	}

	return arrayType
}

func (s *Semantic) analyzeFunctionCall(call *ast.Call) ctypes.Type {
	possibleFuncType := s.analyzeExpression(call.Left)

	if funcType, ok := possibleFuncType.(*ctypes.Function); !ok {
		log.Println(call.Left, s.definedTypes)
		s.error("can't call non function "+call.Left.String()+" of type "+possibleFuncType.String(), call.Token)
	} else {
		if len(call.Parameters) != len(funcType.Parameters) {
			s.error("mismatch number of parameters", call.Token)
		}

		for i, param := range call.Parameters {
			paramType := s.analyzeExpression(param)
			if !s.areTypesEqual(funcType.Parameters[i], paramType) {
				s.typeMismatchError(param.String(), call.Token, funcType.Parameters[i], paramType)
			}
		}

		call.Type = s.unwrapAnonymous(funcType.Return)
		return call.Type
	}

	return ctypes.TODO()
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

	if op == ops.Dot {
		return s.analyzeStructAccess(binaryOperation)
	}

	s.error("can't analyze operator", binaryOperation.Token)
	return ctypes.TODO()
}

func (s *Semantic) analyzeStructAccess(binaryOperation *ast.BinaryOperation) ctypes.Type {
	left := s.analyzeExpression(binaryOperation.Left)
	var strukt *ctypes.Struct
	var isStruct bool

	if ptr, isPointer := left.(*ctypes.Pointer); isPointer {
		strukt, isStruct = s.unwrapAnonymous(ptr.Inner).(*ctypes.Struct)
		if !isStruct {
			s.error("expected struct on access, got "+ptr.Inner.String(), binaryOperation.Token)
			return ctypes.TODO()
		}
	} else {
		strukt, isStruct = left.(*ctypes.Struct)
		if !isStruct {
			s.error("expected struct on access, got "+left.String(), binaryOperation.Token)
			return ctypes.TODO()
		}
	}

	identifier, ok := binaryOperation.Right.(*ast.Identifier)
	if !ok {
		s.error("expected identifier for struct access, got "+binaryOperation.Right.String(), binaryOperation.Token)
		return ctypes.TODO()
	}

	idx, t := strukt.GetField(identifier.Name)
	if idx < 0 || t == nil {
		s.error("unknown struct field "+binaryOperation.String(), binaryOperation.Token)
		return ctypes.TODO()
	}

	return t
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

// replaceAnonymous recursively tries to find an anonymous type and will try to replace it with a true type
func (s *Semantic) replaceAnonymous(t ctypes.Type) ctypes.Type {
	return s.swapTypes(t, ctypes.TODO())
}
