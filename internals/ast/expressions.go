package ast

import (
	"github.com/gabivlj/candice/internals/node"
	"github.com/gabivlj/candice/internals/ops"
	"strconv"
	"strings"
)

type Identifier struct {
	*node.Node
	Name string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string {
	return i.Name
}

type BinaryOperation struct {
	*node.Node
	Left      Expression
	Right     Expression
	Operation ops.Operation
}

func (b *BinaryOperation) expressionNode() {}

func (b *BinaryOperation) String() string {
	return "(" + b.Left.String() + b.Operation.String() + b.Right.String() + ")"
}

type PrefixOperation struct {
	*node.Node
	Right     Expression
	Operation ops.Operation
}

func (p *PrefixOperation) expressionNode() {}

func (p *PrefixOperation) String() string {
	return p.Operation.String() + p.Right.String()
}

type IndexAccess struct {
	*node.Node
	Left   Expression
	Access Expression
}

func (i *IndexAccess) expressionNode() {}

func (i *IndexAccess) String() string {
	return i.Left.String() + "[" + i.Access.String() + "]"
}

// BuiltinCall is a function call that does
// stuff on compile time (like getting the type parameters and generating code accordingly)
type BuiltinCall struct {
	*node.Node
	Name       string
	Parameters []Expression
}

func (bc *BuiltinCall) expressionNode() {}

func (bc *BuiltinCall) String() string {
	builder := strings.Builder{}
	builder.WriteString("@")
	builder.WriteString(bc.Name)
	builder.WriteString("(")
	for i, param := range bc.Parameters {
		if i >= 1 {
			builder.WriteString(", ")
		}
		builder.WriteString(param.String())
	}
	builder.WriteString(")")
	return builder.String()
}

type Call struct {
	*node.Node
	Left       Expression
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
	*node.Node
	Value int64
}

func (i *Integer) expressionNode() {}

func (i *Integer) String() string {
	return strconv.FormatInt(i.Value, 10)
}

type StringLiteral struct {
	*node.Node
	Value string
}

func (s *StringLiteral) String() string { return "\"" + s.Value + "\"" }

func (s *StringLiteral) expressionNode() {}

type StructValue struct {
	Name       string
	Expression Expression
}

type StructLiteral struct {
	*node.Node
	Name   string
	Values []StructValue
}

func (_ *StructLiteral) expressionNode() {}

func (s *StructLiteral) String() string {
	output := strings.Builder{}
	output.WriteString(s.Name)
	output.WriteString("{")
	for _, value := range s.Values {
		output.WriteString(value.Name)
		output.WriteString(": ")
		output.WriteString(value.Expression.String())
		output.WriteString(",\n")
	}
	output.WriteString("}")
	return output.String()
}
