package node

import (
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/pkg/todo"
)

type Node struct {
	Type ctypes.Type
	// TODO This will be filled by the lexer
	Token todo.TodoType
}
