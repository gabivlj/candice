package todo

import "fmt"

// Type is to showcase a variable that its type is not defined yet
type Type = struct{}

// Call is to showcase that a function call should be done here
// but the function is not done yet
func Call(fun string) {
	panic(fmt.Sprintf("function %s is not defined yet", fun))
}
