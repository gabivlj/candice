package ops

import "strconv"

type Operation int

const (
	Plus Operation = iota
	Minus
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

	TempAssign

	// Dot is useful for struct access
	// like struct.thing
	Dot

	// Paren Is not going to be used for
	// ast.BinaryOperation but is useful to retrieve the
	// precedence of a parenthesis or function call.
	Paren
)

var precedences = map[Operation]int{
	AND:              1,
	OR:               1,
	GreaterThan:      2,
	LessThan:         2,
	LessThanEqual:    2,
	GreaterThanEqual: 2,
	Equals:           2,
	NotEquals:        2,
	Plus:             3,
	Minus:            3,
	Multiply:         4,
	Divide:           4,
	BinaryXOR:        5,
	BinaryAND:        5,
	BinaryOR:         5,
	Dot:              6,
	Paren:            6,
}

// Precedence retrieves the precedence of a binary operation, if the precedence
// of the specified operation doesn't apply (maybe your Operation is Bang '!')
// it will return 0
func (o Operation) Precedence() int {
	return precedences[o]
}

func (o Operation) String() string {
	switch o {
	case Plus:
		return "+"
	case Minus:
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
	}

	panic("unknown operand: " + strconv.FormatInt(int64(o), 10))
}
