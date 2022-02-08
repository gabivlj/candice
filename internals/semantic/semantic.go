package semantic

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/node"
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/internals/token"
	"github.com/gabivlj/candice/internals/undomap"
	"github.com/gabivlj/candice/pkg/a"
)

type Semantic struct {
	variables                     *undomap.UndoMap[string, ctypes.Type]
	functionBodies                map[string]*ast.FunctionDeclarationStatement
	definedTypes                  map[string]ctypes.Type
	builtinHandlers               map[string]func(builtin *ast.BuiltinCall) ctypes.Type
	currentStatementBeingAnalyzed ast.Statement
	returns                       bool
	currentExpectedReturnType     ctypes.Type
	Errors                        []error
	insideBreakableBlock          bool
	modules                       map[string]*Semantic
	Root                          *ast.Program
	Compiled                      bool
}

var paths map[string]*Semantic = map[string]*Semantic{}

func (s *Semantic) GetFunction(name string) *ast.FunctionDeclarationStatement {
	return s.functionBodies[name]
}

func (s *Semantic) SizeOf() int64 {
	return 0
}

func (s *Semantic) CandiceType() {}

func (s *Semantic) Alignment() int64 { return 0 }

func (s *Semantic) String() string { return "MODULE" }

func ResetPaths() {
	paths = map[string]*Semantic{}
}

func New() *Semantic {
	s := &Semantic{
		variables:                 undomap.New[string, ctypes.Type](),
		definedTypes:              map[string]ctypes.Type{},
		builtinHandlers:           map[string]func(builtin *ast.BuiltinCall) ctypes.Type{},
		Errors:                    []error{},
		currentExpectedReturnType: ctypes.VoidType,
		returns:                   false,
		modules:                   map[string]*Semantic{},
		functionBodies:            map[string]*ast.FunctionDeclarationStatement{},
	}

	s.builtinHandlers["cast"] = s.analyzeCast
	s.builtinHandlers["alloc"] = s.analyzeAlloc
	s.builtinHandlers["realloc"] = s.analyzeRealloc
	s.builtinHandlers["sizeof"] = s.analyzeSizeOf
	s.builtinHandlers["println"] = s.analyzePrintln
	s.builtinHandlers["free"] = s.analyzeFree

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
	if len(s.Errors) > 1 {
		return
	}
	s.Errors = append(s.Errors, errors.New(fmt.Sprintf("error analyzing on %d:%d (at %s): %s", tok.Line, tok.Position, tok.Type, msg)))
}

func (s *Semantic) typeMismatchError(node string, tok token.Token, expected, got ctypes.Type) {
	message := fmt.Sprintf("_%s_ :: mismatched types, expected=%s, got=%s", node, expected.String(), got.String())
	s.error(message, tok)
}

func (s *Semantic) GetModule(name string) *Semantic {
	return s.modules[name]
}

func (s *Semantic) Analyze(program *ast.Program) {
	s.Root = program
	for _, statement := range program.Statements {
		s.analyzeStatement(statement)
		if len(s.Errors) > 0 {
			return
		}
	}
}

func (s *Semantic) TranslateName(name string) string {
	return ast.CreateIdentifier(ast.RetrieveID(name), s.Root.ID)
}

func (s *Semantic) analyzeStatement(statement ast.Statement) {

	if statement == nil {
		return
	}

	s.currentStatementBeingAnalyzed = statement

	switch statementType := statement.(type) {
	case *ast.ImportStatement:
		s.analyzeImport(statementType)
		return
	case *ast.DeclarationStatement:
		s.analyzeDeclarationStatement(statementType)
		return
	case *ast.StructStatement:
		s.analyzeStructStatement(statementType)
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
	case *ast.ExternStatement:
		s.analyzeExternStatement(statementType)
		return

	case *ast.ExpressionStatement:
		s.analyzeExpression(statementType.Expression)
		return

	case *ast.ReturnStatement:
		s.analyzeReturnStatement(statementType)
		return
	case *ast.GenericTypeDefinition:
		s.analyzeGenericTypeDefinition(statementType)
		return

	case *ast.BreakStatement:
		if !s.insideBreakableBlock {
			s.error("Unexpected break statement", statementType.Token)
		}
		return

	case *ast.TypeDefinition:
		s.analyzeTypeDefinition(statementType)
		return

	case *ast.ContinueStatement:
		if !s.insideBreakableBlock {
			s.error("Unexpected continue statement", statementType.Token)
		}
		return
	}

	log.Fatalln("couldn't analyze statement: " + statement.String() + " ")
}

func (s *Semantic) analyzeTypeDefinition(typeDef *ast.TypeDefinition) {
	s.replaceAnonymous(typeDef.Type)
	trueType := s.UnwrapAnonymous(typeDef.Type)
	typeDef.Type = trueType
	s.definedTypes[typeDef.Name] = trueType
}

func (s *Semantic) analyzeGenericTypeDefinition(genericType *ast.GenericTypeDefinition) {

	return
}

func (s *Semantic) analyzeExternStatement(extern *ast.ExternStatement) {
	funk, ok := extern.Type.(*ctypes.Function)
	if !ok {
		s.typeMismatchError(extern.String(), extern.Token, &ctypes.Function{Name: "function"}, extern.Type)
	}
	s.variables.Add(funk.Name, funk)
}

func (s *Semantic) analyzeAssigmentStatement(assign *ast.AssignmentStatement) {
	right := s.analyzeExpression(assign.Expression)
	left := s.analyzeExpression(assign.Left)
	if !s.areTypesEqual(left, right) {
		s.typeMismatchError(assign.String(), s.currentStatementBeingAnalyzed.GetToken(), left, right)
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

	if forStatement.Condition == nil {
		forStatement.Condition = &ast.Integer{Value: 1, Node: &node.Node{Type: ctypes.I32}}
	}

	condition := s.analyzeExpression(forStatement.Condition)

	if !ctypes.IsNumeric(condition) && condition != ctypes.VoidType {
		s.typeMismatchError(forStatement.Condition.String(), forStatement.Token, ctypes.I32, condition)
	}

	s.analyzeStatement(forStatement.Operation)
	tempInsideBlock := s.insideBreakableBlock
	s.insideBreakableBlock = true
	s.analyzeBlock(forStatement.Block)
	s.insideBreakableBlock = tempInsideBlock
	s.leaveFrame()
}

func (s *Semantic) analyzeFunctionStatement(fun *ast.FunctionDeclarationStatement) {
	if fun.FunctionType.Return == nil {
		fun.FunctionType.Return = ctypes.VoidType
	}

	s.variables.Add(fun.FunctionType.Name, fun.FunctionType)
	s.functionBodies[fun.FunctionType.Name] = fun

	s.enterFrame()

	for i, param := range fun.FunctionType.Parameters {
		// try to replace the anonymous type with its true type
		fun.FunctionType.Parameters[i] = s.replaceAnonymous(param)
		s.variables.Add(fun.FunctionType.Names[i], fun.FunctionType.Parameters[i])
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
	fun.FunctionType.Return = s.UnwrapAnonymous(fun.FunctionType.Return)
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
			s.typeMismatchError(currentIf.Condition.String(), currentIf.GetToken(), ctypes.I32, condition)
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
	theType := s.UnwrapAnonymous(s.analyzeExpression(returnStatement.Expression))
	if !s.areTypesEqual(theType, s.currentExpectedReturnType) {
		s.typeMismatchError(returnStatement.String(), returnStatement.Token, s.currentExpectedReturnType, theType)
	}
	returnStatement.Type = s.UnwrapAnonymous(theType)
	s.returns = true
}

func (s *Semantic) analyzeStructStatement(statementType *ast.StructStatement) {

	s.definedTypes[statementType.Type.Name] = statementType.Type
	for _, t := range statementType.Type.Fields {
		unwrappedType := s.unwrap(t)
		if anonymous, ok := unwrappedType.(*ctypes.Anonymous); ok {
			definedType := s.UnwrapAnonymous(anonymous)
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
		t := builtinHandler(call)
		call.Type = t
		return t
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

func (s *Semantic) UnwrapAnonymous(t ctypes.Type) ctypes.Type {
	if t == ctypes.TODO() {
		return t
	}

	if anonymous, ok := t.(*ctypes.Anonymous); ok {
		module := ""
		if anonymous.Modules != nil && len(anonymous.Modules) > 0 {
			module = anonymous.Modules[0]
		}

		semantic := s.retrieveModule(module)
		name := semantic.TranslateName(anonymous.Name)
		// replace anonymous type name to the module one.
		anonymous.Name = name

		t, ok := semantic.definedTypes[name]
		if !ok {
			typesDefined := ""
			for _, t := range semantic.definedTypes {
				typesDefined += "- " + t.String() + "\n"
			}

			if typesDefined == "" {
				typesDefined = "No types defined in the module."
			}

			s.error("Couldn't guess type "+ast.RetrieveID(anonymous.Name)+", maybe spelt the type wrong? These are the defined types in the module"+" "+ast.RetrieveID(module)+":\n"+typesDefined, s.currentStatementBeingAnalyzed.GetToken())
			return ctypes.TODO()
		}

		return t
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

	if _, stopRecursionForStruct := t.(*ctypes.Struct); stopRecursionForStruct {
		return t
	}

	if toSwap == ctypes.TODO() {
		trueType := s.UnwrapAnonymous(t)
		if _, ok := trueType.(*ctypes.Anonymous); (ok && trueType == t) || trueType == nil {
			s.error("unknown type "+t.String(), s.currentStatementBeingAnalyzed.GetToken())
		}
		return trueType
	}

	return toSwap
}

func (s *Semantic) areTypesEqual(first, second ctypes.Type) bool {
	first = s.UnwrapAnonymous(first)
	second = s.UnwrapAnonymous(second)

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
	case *ast.Float:
		return s.analyzeFloat(expressionType)
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
	case *ast.StringLiteral:
		stringLiteralType := &ctypes.Pointer{Inner: ctypes.I8}
		expressionType.Type = stringLiteralType
		return stringLiteralType

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
		indexAccess.Type = arr.Inner
		return arr.Inner
	}

	if ptr, ok := leftType.(*ctypes.Pointer); ok {
		indexAccess.Type = ptr.Inner
		return ptr.Inner
	}

	s.error("mismatched types on index access, internal compiler bug", indexAccess.Token)

	return ctypes.TODO()
}

func (s *Semantic) retrieveModule(moduleName string) *Semantic {
	if moduleName == "" {
		return s
	}

	module, ok := s.modules[moduleName]

	if !ok {
		s.Errors = append(s.Errors, errors.New("undefined module "+moduleName))
		return s
	}

	return module
}

func (s *Semantic) retrieveTypeFromStruct(structLiteral *ast.StructLiteral) (ctypes.Type, error) {
	module := s.retrieveModule(structLiteral.Module)
	// TODO: change here to translate into different code
	structLiteral.Name = module.TranslateName(structLiteral.Name)
	structType, ok := module.definedTypes[structLiteral.Name]

	if !ok {
		return nil, errors.New("undefined struct " + structLiteral.Name + ": " + structLiteral.String())
	}

	return structType, nil
}

func (s *Semantic) analyzeStructLiteral(structLiteral *ast.StructLiteral) ctypes.Type {

	possibleStructType, err := s.retrieveTypeFromStruct(structLiteral)

	if err != nil {
		s.Errors = append(s.Errors, err)

		return ctypes.TODO()
	}

	structType, ok := s.UnwrapAnonymous(possibleStructType).(*ctypes.Struct)

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

	arrayLiteral.Type = arrayType
	return arrayType
}

func (s *Semantic) analyzeFunctionCall(call *ast.Call) ctypes.Type {
	possibleFuncType := s.analyzeExpression(call.Left)

	if funcType, ok := possibleFuncType.(*ctypes.Function); !ok {
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

		call.Type = funcType.Return
		return call.Type
	}

	return ctypes.TODO()
}

func (s *Semantic) analyzeSimpleIdentifier(identifier *ast.Identifier) ctypes.Type {
	if module, ok := s.modules[identifier.Name]; ok {
		identifier.Type = module
		return module
	}

	if identifierType := s.variables.Get(identifier.Name); identifierType != nil {
		identifier.Type = identifierType
		return identifierType
	}
	s.error("undefined variable "+identifier.Name, identifier.Token)
	return ctypes.TODO()
}

func (s *Semantic) analyzePrefixOperation(prefixOperation *ast.PrefixOperation) ctypes.Type {
	t := s.UnwrapAnonymous(s.analyzeExpression(prefixOperation.Right))
	prefixOperation.Type = t
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
		prefixOperation.Type = &ctypes.Pointer{Inner: t}
		return prefixOperation.Type
	}

	if prefixOperation.Operation == ops.Multiply {
		if ptr, ok := t.(*ctypes.Pointer); !ok {
			s.typeMismatchError(prefixOperation.String(), prefixOperation.Token, &ctypes.Pointer{Inner: t}, t)
			return t
		} else {
			prefixOperation.Type = s.UnwrapAnonymous(ptr.Inner)
			return prefixOperation.Type
		}
	}

	s.error("unknown prefix operator to analyze", prefixOperation.Token)

	return t
}

func (s *Semantic) analyzeBinaryOperation(binaryOperation *ast.BinaryOperation) ctypes.Type {
	op := binaryOperation.Operation
	if s.isArithmetic(op) {
		t := s.analyzeArithmetic(binaryOperation)
		binaryOperation.Type = t
		return t
	}

	if op == ops.Dot {
		t := s.analyzeStructAccess(binaryOperation)
		binaryOperation.Type = t
		return t
	}

	s.error("can't analyze operator", binaryOperation.Token)
	return ctypes.TODO()
}

func (s *Semantic) analyzeModuleAccess(module *Semantic, binaryOp *ast.BinaryOperation) ctypes.Type {
	identifier, ok := binaryOp.Right.(*ast.Identifier)
	if !ok {
		s.error("expected identifier for module access, got "+binaryOp.Right.String(), binaryOp.Token)
		return ctypes.TODO()
	}

	name := module.TranslateName(identifier.Name)
	accessedElement := module.variables.Get(name)

	// Reassign identifier to the new name
	identifier.Name = name
	if accessedElement == nil {
		s.error(ast.RetrieveID(identifier.Name)+" does not exist in the specified module", binaryOp.Token)
		return ctypes.TODO()
	}

	binaryOp.Type = accessedElement

	return accessedElement
}

func (s *Semantic) analyzeStructAccess(binaryOperation *ast.BinaryOperation) ctypes.Type {
	left := s.analyzeExpression(binaryOperation.Left)
	var strukt *ctypes.Struct
	var isStruct bool

	if module, isModule := left.(*Semantic); isModule {
		return s.analyzeModuleAccess(module, binaryOperation)
	}

	if ptr, isPointer := left.(*ctypes.Pointer); isPointer {
		strukt, isStruct = s.UnwrapAnonymous(ptr.Inner).(*ctypes.Struct)
		if !isStruct {
			s.error("expected struct on access, got "+ptr.Inner.String(), binaryOperation.Token)
			return ctypes.TODO()
		}
	} else {
		strukt, isStruct = s.UnwrapAnonymous(left).(*ctypes.Struct)
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

	// Just in case that the identifier.Name is poisoned by ID generation, extract original name
	identifier.Name = ast.RetrieveID(identifier.Name)
	idx, t := strukt.GetField(identifier.Name)
	if idx < 0 || t == nil {
		s.error("unknown struct field "+binaryOperation.String(), binaryOperation.Token)
		return ctypes.TODO()
	}

	binaryOperation.Type = t
	return t
}

func (s *Semantic) isArithmetic(op ops.Operation) bool {
	return op == ops.OR || op == ops.Multiply || op == ops.BinaryXOR || op == ops.BinaryOR ||
		op == ops.BinaryAND || op == ops.AND || op == ops.Add || op == ops.Subtract || op == ops.LessThanEqual ||
		op == ops.LessThan || op == ops.Equals || op == ops.GreaterThan || op == ops.GreaterThanEqual ||
		op == ops.NotEquals || op == ops.Divide
}

// analyzeImport works in a really tricky way
// - First off we start by gathering all types that the file is gonna need as parameters
// - Then we should parse and analyze the file that is being imported if it hasn't been yet
// - Now this is a tricky part, we check if the file with those type parameters have been analyzed before
//  and if it does we try to get the semantic component and put it available as that module name.
// - Because they are the exact same modules when using them in 2 different parts of the project,
//   we won't find type discrepancies, which is good because really, it's the same type.
// - This works because let's remember that string names in definitions are <name_put_by_the_user> '-' <random_id_set_by_the_parser>.
//  	the random id is located in the *Semantic.Root attribute, so we can use it to create or parse names.
func (s *Semantic) analyzeImport(importStatement *ast.ImportStatement) {
	types := make([]ctypes.Type, 0, len(importStatement.Types))
	for _, t := range importStatement.Types {
		t = s.replaceAnonymous(t)
		types = append(types, s.UnwrapAnonymous(t))
	}

	// TODO this is not correct, should be relative path of current file
	currentPath, err := os.Getwd()
	if err != nil {
		s.Errors = append(s.Errors, err)
		currentPath = "/"
	}

	path := path.Join(currentPath, importStatement.Path.Value)
	hash := strings.Builder{}

	for _, t := range types {
		hash.WriteByte(',')
		hash.WriteString(t.String())
	}
	endHash := path + hash.String()

	if existingSemantic, ok := paths[endHash]; ok {
		s.modules[importStatement.Name] = existingSemantic
		return
	}

	text, err := os.ReadFile(path)
	if err != nil {
		s.error(fmt.Sprintf("error importing file with path %s: %s", importStatement.Path, err.Error()), importStatement.Token)
		return
	}

	l := lexer.New(string(text))
	p := parser.New(l)
	p.TypeParameters = types
	tree := p.Parse()
	if len(p.Errors) > 0 {
		s.error("error parsing file imported on path "+importStatement.Path.String(), importStatement.Token)
		s.Errors = append(s.Errors, p.Errors...)
		return
	}

	internalSemantic := New()
	internalSemantic.Analyze(tree)
	if len(internalSemantic.Errors) > 0 {
		s.error("error analyzing file imported on path "+importStatement.Path.String(), importStatement.Token)
		s.Errors = append(s.Errors, internalSemantic.Errors...)
		return
	}

	paths[endHash] = internalSemantic

	s.modules[importStatement.Name] = internalSemantic
}

func (s *Semantic) analyzeArithmetic(binaryOperation *ast.BinaryOperation) ctypes.Type {
	left := s.UnwrapAnonymous(s.analyzeExpression(binaryOperation.Left))
	right := s.UnwrapAnonymous(s.analyzeExpression(binaryOperation.Right))
	if !ctypes.IsNumeric(left) {
		s.error("expected numeric type, got: "+left.String(), binaryOperation.Token)
	}

	if !ctypes.IsNumeric(right) {
		s.error("expected numeric type, got: "+right.String(), binaryOperation.Token)
	}

	if !s.areTypesEqual(left, right) {
		s.typeMismatchError(binaryOperation.String(), binaryOperation.Token, right, left)
	}

	if binaryOperation.Operation.IsComparison() {
		return ctypes.I1
	}

	return left
}

func (s *Semantic) analyzeInteger(integer *ast.Integer) ctypes.Type {
	return integer.Type
}

func (s *Semantic) analyzeFloat(float *ast.Float) ctypes.Type {
	return float.Type
}

// replaceAnonymous recursively tries to find an anonymous type and will try to replace it with a true type
func (s *Semantic) replaceAnonymous(t ctypes.Type) ctypes.Type {
	return s.swapTypes(t, ctypes.TODO())
}
