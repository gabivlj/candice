package ops

import "strconv"

type Operation int

const (
	Add Operation = iota
	Subtract
	Multiply
	Divide
	AND
	OR
	At
	Bang
	BinaryXOR
	BinaryAND
	BinaryOR
	Reference
	GreaterThan
	LessThan
	GreaterThanEqual
	LessThanEqual
	Equals
	NotEquals
	Assigns
	AddOne
	TempAssign

	// Dot is useful for struct access
	// like struct.thing
	Dot

	// Paren Is not going to be used for
	// ast.BinaryOperation but is useful to retrieve the
	// precedence of a parenthesis or function call.
	Paren

	As
)

func (o Operation) String() string {
	switch o {
	case Add:
		return "+"
	case Subtract:
		return "-"
	case Multiply:
		return "*"
	case Divide:
		return "/"
	case AND:
		return "&&"
	case OR:
		return "||"
	case Bang:
		return "!"
	case BinaryXOR:
		return "^"
	case BinaryAND:
		return "&"
	case BinaryOR:
		return "|"
	case Reference:
		return "&"
	case GreaterThan:
		return ">"
	case GreaterThanEqual:
		return ">="
	case LessThan:
		return "<"
	case LessThanEqual:
		return "<="
	case Equals:
		return "=="
	case Assigns:
		return "="
	case NotEquals:
		return "!="
	case Paren:
		return "("
	case Dot:
		return "."
	case As:
		return "as"
	}

	panic("unknown operand: " + strconv.FormatInt(int64(o), 10))
}

func (o Operation) IsComparison() bool {
	return o == GreaterThanEqual || o == Equals || o == NotEquals || o == GreaterThan ||
		o == LessThanEqual || o == LessThan || o == AND || o == OR
}
