package ast

import (
	"strings"

	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/token"
)

type Node interface {
	String() string
	GetToken() token.Token
}

type Expression interface {
	Node
	expressionNode()
	GetType() ctypes.Type
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	ID         string
	Statements []Statement
}

func (p *Program) GetToken() token.Token {
	if len(p.Statements) == 0 {
		return token.Token{}
	}

	return p.Statements[0].GetToken()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) GetToken() token.Token {
	return e.Token
}

func (e *ExpressionStatement) String() string {
	return e.Expression.String() + ";"
}

func (e *ExpressionStatement) statementNode() {}

func (p *Program) String() string {
	builder := strings.Builder{}
	for _, s := range p.Statements {
		builder.WriteString(s.String() + "\n")
	}
	return builder.String()
}

func CreateIdentifier(name string, id string) string {
	if name == "main" {
		return name
	}
	return name + "-" + id
}

func RetrieveID(name string) string {
	return strings.Split(name, "-")[0]
}

func RetrieveRightID(name string) string {
	return strings.Split(name, "-")[1]
}
