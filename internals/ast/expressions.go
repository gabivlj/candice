package ast

import (
	"github.com/gabivlj/candice/internals/ops"
	"strconv"
	"strings"
)

type Identifier struct {
	Name string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string {
	return i.Name
}

type BinaryOperation struct {
	Left Expression
	Right Expression
	Operation ops.Operation
}

func (b *BinaryOperation) expressionNode() {}

func (b *BinaryOperation) String() string {
	return "("+b.Left.String() + b.Operation.String() + b.Right.String()+")"
}

type PrefixOperation struct {
	Right Expression
	Operation ops.Operation
}

func (p *PrefixOperation) expressionNode() {}

func (p *PrefixOperation) String() string {
	return p.Operation.String() + p.Right.String()
}

type IndexAccess struct {
	Left Expression
	Access Expression
}

func (i *IndexAccess) expressionNode() {}

func (i *IndexAccess) String() string {
	return i.Left.String() + "[" + i.Access.String() + "]"
}

type Call struct {
	Left Expression
	Parameters []Expression
}

func (c *Call) expressionNode() {}

func (c *Call) String() string {
	builder := strings.Builder{}
	builder.WriteString(c.Left.String())
	builder.WriteString("(")
	for i, param := range c.Parameters {
		if i >= 1 {
			builder.WriteString(", ")
		}
		builder.WriteString(param.String())
	}
	builder.WriteString(")")
	return builder.String()
}

type Integer struct {
	Value int64
}

func (i *Integer) expressionNode() {}

func (i *Integer) String() string {
	return strconv.FormatInt(i.Value, 10)
}

type StringLiteral struct {
	Value string
}

func (s *StringLiteral) String() string { return "\"" + s.Value + "\"" }

func (s *StringLiteral) expressionNode() {}

