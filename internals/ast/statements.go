package ast

import (
	"fmt"
	"github.com/gabivlj/candice/internals/ctypes"
	"strings"
)

type Block struct {
	Statements []Statement
}

type ConditionPlusBlock struct {
	Block *Block
	Condition Expression
}

func (c *ConditionPlusBlock) statementNode(){}

func (c *ConditionPlusBlock) String() string {
	s := strings.Builder{}
	s.WriteString(c.Condition.String())
	s.WriteString(" {\n")
	s.WriteString(c.Block.String())
	s.WriteString("\n}")
	return s.String()
}

func (b *Block) String() string {
	s := strings.Builder{}
	for i, statement := range b.Statements {
		if i >= 1 {
			s.WriteByte('\n')
		}
		s.WriteString(statement.String())
	}
	return s.String()
}

type StructStatement struct {
	Type *ctypes.Struct
}

func (s *StructStatement) statementNode() {}

func (s *StructStatement) String() string {
	return s.Type.FullString()
}

type DeclarationStatement struct {
	Name       string
	Type       ctypes.Type
	Expression Expression
}

func (_ *DeclarationStatement) statementNode() {}

func (d *DeclarationStatement) String() string {
	return fmt.Sprintf("%s :%s = %s", d.Name, d.Type.String(), d.Expression.String())
}

type IfStatement struct {
	Condition Expression
	Block     *Block
	ElseIfs   []*ConditionPlusBlock
	Else      *Block
}

func (i *IfStatement) String() string {
	s := strings.Builder{}
	s.WriteString("if ")
	s.WriteString(i.Condition.String())
	s.WriteString(" {")
	if i.Block != nil {
		s.WriteString("\n")
		s.WriteString(i.Block.String())
		s.WriteString("\n")
	}
	s.WriteString("}")

	for _, iff := range i.ElseIfs {
		s.WriteString(" else if ")
		s.WriteString(iff.String())
	}

	if i.Else != nil {
		s.WriteString(" else ")
		s.WriteString("{\n")
		s.WriteString(i.Else.String())
		s.WriteString("\n}")
	}

	return s.String()
}