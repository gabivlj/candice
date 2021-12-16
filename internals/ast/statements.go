package ast

import (
	"fmt"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/token"
	"strings"
)

type Block struct {
	Statements []Statement
}

type ConditionPlusBlock struct {
	Block     *Block
	Condition Expression
}

func (c *ConditionPlusBlock) statementNode() {}

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
	Token token.Token
	Type  *ctypes.Struct
}

func (s *StructStatement) statementNode() {}

func (s *StructStatement) String() string {
	return s.Type.FullString()
}

type DeclarationStatement struct {
	Token      token.Token
	Name       string
	Type       ctypes.Type
	Expression Expression
}

func (_ *DeclarationStatement) statementNode() {}

func (d *DeclarationStatement) String() string {
	return fmt.Sprintf("%s :%s = %s;", d.Name, d.Type.String(), d.Expression.String())
}

type AssignmentStatement struct {
	Left       Expression
	Expression Expression
}

func (_ *AssignmentStatement) statementNode() {}

func (d *AssignmentStatement) String() string {
	return fmt.Sprintf("%s = %s;", d.Left.String(), d.Expression.String())
}

type ForStatement struct {
	Token                token.Token
	Condition            Expression
	InitializerStatement Statement
	Operation            Statement
	Block                *Block
}

func (f *ForStatement) statementNode() {}

func (f *ForStatement) String() string {
	s := strings.Builder{}
	s.WriteString("for")

	if f.InitializerStatement != nil {
		s.WriteByte(' ')
		s.WriteString(f.InitializerStatement.String())
	}

	if f.Condition != nil {
		s.WriteByte(' ')
		s.WriteString(f.Condition.String())
	}

	if f.Operation != nil {
		s.WriteString("; ")
		s.WriteString(f.Operation.String())
	} else if f.InitializerStatement != nil {
		s.WriteString(";")
	}

	s.WriteString(" {\n")
	s.WriteString(f.Block.String())
	s.WriteString("\n}")
	return s.String()
}

type IfStatement struct {
	Token     token.Token
	Condition Expression
	Block     *Block
	ElseIfs   []*ConditionPlusBlock
	Else      *Block
}

func (i *IfStatement) statementNode() {}

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
