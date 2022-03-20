package parser

type BuiltinFunctionParseRequirements struct {
	Types      int
	Parameters int
}

var builtinFunctions = map[string]BuiltinFunctionParseRequirements{}

const UndefinedNumberOfParameters = -1

// AddBuiltinFunction adds a requirement of parsing a builtin function
func (p *Parser) addBuiltinFunction(name string, numberOfTypes, numberOfParameters int) {
	builtinFunctions[name] = BuiltinFunctionParseRequirements{
		Types:      numberOfTypes,
		Parameters: numberOfParameters,
	}
}

func (p *Parser) getBuiltinFunctionRequirements(name string) BuiltinFunctionParseRequirements {
	return builtinFunctions[name]
}

func (p *Parser) initBuiltinFunctions() {
	p.addBuiltinFunction("alloc", 1, 1)
	p.addBuiltinFunction("print", 0, UndefinedNumberOfParameters)
	p.addBuiltinFunction("cast", 1, 1)
	p.addBuiltinFunction("realloc", 0, 2)
	p.addBuiltinFunction("free", 0, 1)
	p.addBuiltinFunction("sizeof", 1, 0)
	p.addBuiltinFunction("unreachable", 0, 0)
	p.addBuiltinFunction("asm", 1, UndefinedNumberOfParameters)
}
