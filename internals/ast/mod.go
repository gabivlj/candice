package ast

import (
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/token"
	"strings"
)

type Node interface {
	String() string
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

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
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
