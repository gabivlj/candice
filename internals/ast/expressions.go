package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/node"
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/internals/token"
	"github.com/gabivlj/candice/pkg/split"
)

type Identifier struct {
	*node.Node
	Name string
}

func (i *Identifier) GetType() ctypes.Type {
	return i.Node.Type
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string {
	return RetrieveID(i.Name)
}

func (i *Identifier) GetToken() token.Token {
	return i.Token
}

type BinaryOperation struct {
	*node.Node
	Left      Expression
	Right     Expression
	Operation ops.Operation
}

func (b *BinaryOperation) GetType() ctypes.Type {
	return b.Node.Type
}

func (b *BinaryOperation) expressionNode() {}

func (b *BinaryOperation) GetToken() token.Token {
	return b.Token
}

func (b *BinaryOperation) String() string {
	if b.Operation == ops.Dot {
		return b.Left.String() + b.Operation.String() + b.Right.String()
	}
	return "(" + b.Left.String() + " " + b.Operation.String() + " " + b.Right.String() + ")"
}

type PrefixOperation struct {
	*node.Node
	Right     Expression
	Operation ops.Operation
}

func (p *PrefixOperation) GetType() ctypes.Type {
	return p.Node.Type
}

func (p *PrefixOperation) expressionNode() {}

func (p *PrefixOperation) GetToken() token.Token {
	return p.Token
}

func (p *PrefixOperation) String() string {
	return p.Operation.String() + p.Right.String()
}

type IndexAccess struct {
	*node.Node
	Left   Expression
	Access Expression
}

func (i *IndexAccess) GetType() ctypes.Type {
	return i.Node.Type
}

func (i *IndexAccess) expressionNode() {}

func (i *IndexAccess) GetToken() token.Token {
	return i.Token
}

func (i *IndexAccess) String() string {
	return i.Left.String() + "[" + i.Access.String() + "]"
}

// BuiltinCall is a function call that does
// stuff on compile time (like getting the type parameters and generating code accordingly)
type BuiltinCall struct {
	*node.Node
	Name           string
	TypeParameters []ctypes.Type
	Parameters     []Expression
}

func (bc *BuiltinCall) GetType() ctypes.Type {
	return bc.Node.Type
}

func (bc *BuiltinCall) expressionNode() {}

func (bc *BuiltinCall) GetToken() token.Token {
	return bc.Token
}

func (bc *BuiltinCall) castString() string {
	return fmt.Sprintf("%s as %s", bc.Parameters[0], bc.TypeParameters[0])
}

func (bc *BuiltinCall) String() string {
	if bc.Name == "cast" {
		return bc.castString()
	}

	builder := strings.Builder{}
	builder.WriteString("@")
	builder.WriteString(bc.Name)
	builder.WriteString("(")
	for i, param := range bc.TypeParameters {
		if i >= 1 {
			builder.WriteString(", ")
		}
		builder.WriteString(param.String())
	}

	for i, param := range bc.Parameters {
		if i >= 1 || len(bc.TypeParameters) > 0 {
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

func (c *Call) GetType() ctypes.Type {
	return c.Node.Type
}

func (c *Call) expressionNode() {}

func (c *Call) GetToken() token.Token {
	return c.Token
}

func (c *Call) String() string {
	builder := strings.Builder{}
	builder.WriteString(RetrieveID(c.Left.String()))
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

func (i *Integer) GetType() ctypes.Type {
	return i.Node.Type
}

func (i *Integer) expressionNode() {}

func (i *Integer) GetToken() token.Token {
	return i.Token
}

func (i *Integer) String() string {
	if i.Node == nil {
		return strconv.FormatInt(i.Value, 10)
	}

	if i.Token.Type == token.CHAR {
		return fmt.Sprintf("'%c'", i.Value)
	}

	return strconv.FormatInt(i.Value, 10)
}

type Float struct {
	*node.Node
	Value float64
}

func (i *Float) GetType() ctypes.Type {
	return i.Node.Type
}

func (i *Float) expressionNode() {}

func (i *Float) GetToken() token.Token {
	return i.Token
}

func (i *Float) String() string {
	return strconv.FormatFloat(i.Value, 'f', -1, 64)
}

type StringLiteral struct {
	*node.Node
	Value string
}

func (s *StringLiteral) GetType() ctypes.Type {
	return s.Node.Type
}

func (s *StringLiteral) String() string { return "\"" + s.Value + "\"" }

func (s *StringLiteral) expressionNode() {}

func (s *StringLiteral) GetToken() token.Token {
	return s.Token
}

type StructValue struct {
	Name       string
	Expression Expression
}

type StructLiteral struct {
	*node.Node
	Name   string
	Values []StructValue
	Module string
}

func (s *StructLiteral) GetType() ctypes.Type {
	return s.Node.Type
}

func (_ *StructLiteral) expressionNode() {}

func (s *StructLiteral) GetToken() token.Token {
	return s.Token
}

func (s *StructLiteral) String() string {
	output := strings.Builder{}
	output.WriteByte('@')
	output.WriteString(RetrieveID(s.Name))
	output.WriteString("{\n")
	for _, value := range s.Values {
		output.WriteString(value.Name)
		output.WriteString(": ")
		output.WriteString(value.Expression.String())
		output.WriteString(",\n")
	}
	output.WriteString("}")
	return output.String()
}

type ArrayLiteral struct {
	*node.Node
	Values []Expression
}

func (a *ArrayLiteral) GetType() ctypes.Type {
	return a.Node.Type
}

func (a *ArrayLiteral) expressionNode() {}

func (a *ArrayLiteral) GetToken() token.Token {
	return a.Token
}

func (a *ArrayLiteral) String() string {
	builder := strings.Builder{}
	builder.WriteString(a.Node.Type.String())
	builder.WriteString(" {")
	var expressions []string
	for _, expr := range a.Values {
		expressions = append(expressions, expr.String())
	}
	builder.WriteString(strings.Join(expressions, ", "))
	builder.WriteString("}")
	return builder.String()
}

type AnonymousFunction struct {
	Token        token.Token
	FunctionType *ctypes.Function
	Block        *Block
}

func (f *AnonymousFunction) GetType() ctypes.Type {
	return f.FunctionType
}

func (f *AnonymousFunction) expressionNode() {}

func (f *AnonymousFunction) String() string {
	builder := strings.Builder{}
	builder.WriteString(f.FunctionType.FullString())
	builder.WriteString(" {\n")
	builder.WriteString(f.Block.String())
	builder.WriteString("\n}")
	return builder.String()
}

func (f *AnonymousFunction) GetToken() token.Token { return f.Token }

// CommaExpressions are expressions separated by commas
type CommaExpressions struct {
	*node.Node
	Token       token.Token
	Expressions []Expression
}

func (c *CommaExpressions) expressionNode() {}
func (c *CommaExpressions) String() string {
	return split.Split(c.Expressions, ", ")
}
func (c *CommaExpressions) GetType() ctypes.Type  { return c.Node.Type }
func (c *CommaExpressions) GetToken() token.Token { return c.GetToken() }
