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
	Assigns
	Dot
)

func (o Operation) String() string {
	switch o {
	case Plus: return "+"
	case Minus: return "-"
	case Multiply: return "*"
	case Divide: return "/"
	case AND: return "&&"
	case OR: return "||"
	case Bang: return "!"
	case BinaryXOR: return "^"
	case BinaryAND: return "&"
	case BinaryOR: return "|"
	case Reference: return "&"
	case GreaterThan: return ">"
	case GreaterThanEqual: return ">="
	case LessThan: return "<"
	case LessThanEqual: return "<="
	case Equals: return "=="
	case Assigns: return "="
	case Dot: return "."
	}

	panic("unknown operand: " + strconv.FormatInt(int64(o), 10))
}
