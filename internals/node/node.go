package node

import "github.com/gabivlj/candice/pkg/todo"

type Node struct {
	// TODO: This in the future will be filled by the semantic analyzer component
	SemanticType todo.TodoType
	// TODO This will be filled by the lexer
	Token todo.TodoType
}
