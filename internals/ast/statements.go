package ast

import (
	"fmt"
	"strings"

	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/token"
)

type Block struct {
	Statements []Statement
	Token      token.Token
}

type MacroBlock struct {
	*Block
}

func (b *Block) GetToken() token.Token {
	return b.Token
}

func (b *Block) statementNode() {}

type ConditionPlusBlock struct {
	Block     *Block
	Condition Expression
}

func (c *ConditionPlusBlock) statementNode() {}

func (c *ConditionPlusBlock) GetToken() token.Token {
	return token.Token{}
}

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

func (s *StructStatement) GetToken() token.Token {
	return s.Token
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

func (d *DeclarationStatement) GetToken() token.Token {
	return d.Token
}

func (_ *DeclarationStatement) statementNode() {}

func (d *DeclarationStatement) String() string {
	return fmt.Sprintf("%s :%s = %s;", RetrieveID(d.Name), d.Type.String(), d.Expression.String())
}

type AssignmentStatement struct {
	Token      token.Token
	Left       Expression
	Expression Expression
}

func (_ *AssignmentStatement) statementNode() {}

func (d *AssignmentStatement) String() string {
	return fmt.Sprintf("%s = %s;", d.Left.String(), d.Expression.String())
}

func (d *AssignmentStatement) GetToken() token.Token { return d.Token }

type ForStatement struct {
	Token                token.Token
	Condition            Expression
	InitializerStatement Statement
	Operation            Statement
	Block                *Block
}

func (f *ForStatement) statementNode() {}

func (f *ForStatement) GetToken() token.Token { return f.Token }

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

func (i *IfStatement) GetToken() token.Token { return i.Token }

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

type FunctionDeclarationStatement struct {
	Token        token.Token
	FunctionType *ctypes.Function
	Block        *Block
}

func (f *FunctionDeclarationStatement) statementNode() {}

func (f *FunctionDeclarationStatement) String() string {
	builder := strings.Builder{}
	builder.WriteString(f.FunctionType.FullString())
	builder.WriteString(" {\n")
	builder.WriteString(f.Block.String())
	builder.WriteString("\n}")
	return builder.String()
}

func (f *FunctionDeclarationStatement) GetToken() token.Token { return f.Token }

type ExternStatement struct {
	Token token.Token
	Type  ctypes.Type
}

func (e *ExternStatement) String() string {
	return "extern " + e.Type.String() + ";"
}

func (e *ExternStatement) statementNode() {}

func (e *ExternStatement) GetToken() token.Token { return e.Token }

type ReturnStatement struct {
	Type       ctypes.Type
	Token      token.Token
	Expression Expression
}

func (r *ReturnStatement) String() string {
	returnExpr := ""
	if r.Expression != nil {
		returnExpr = r.Expression.String()
	}
	return "return " + returnExpr + ";"
}

func (r *ReturnStatement) statementNode() {}

func (r *ReturnStatement) GetToken() token.Token { return r.Token }

type ImportStatement struct {
	Token token.Token
	Name  string
	Types []ctypes.Type
	Path  *StringLiteral
}

func (_ *ImportStatement) statementNode() {}

func (i *ImportStatement) String() string {
	var types []string
	for _, t := range i.Types {
		types = append(types, t.String())
	}
	return "import " + RetrieveID(i.Name) + ", " + strings.Join(types, ", ") + ", " + i.Path.String()
}

func (i *ImportStatement) GetToken() token.Token { return i.Token }

type BreakStatement struct {
	Token token.Token
}

func (b *BreakStatement) statementNode() {}

func (b *BreakStatement) String() string { return b.Token.Literal }

func (b *BreakStatement) GetToken() token.Token { return b.Token }

type ContinueStatement struct {
	Token token.Token
}

func (c *ContinueStatement) statementNode() {}

func (c *ContinueStatement) String() string        { return c.Token.Literal }
func (c *ContinueStatement) GetToken() token.Token { return c.Token }

type GenericTypeDefinition struct {
	Token        token.Token
	Name         string
	ReplacedType ctypes.Type
}

func (g *GenericTypeDefinition) statementNode() {}

func (g *GenericTypeDefinition) String() string        { return g.Name }
func (g *GenericTypeDefinition) GetToken() token.Token { return g.Token }

type TypeDefinition struct {
	Token token.Token
	Name  string
	Type  ctypes.Type
}

func (g *TypeDefinition) statementNode() {}

func (g *TypeDefinition) String() string { return fmt.Sprintf("type %s = %s", g.Name, g.Type.String()) }

func (g *TypeDefinition) GetToken() token.Token { return g.Token }
