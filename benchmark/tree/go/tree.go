package main

func runTreeBench() {

	for i := 0; i < 500; i++ {

		c := NewTree(20)
		c.Count()

	}
}

type Tree struct {
	Left  *Tree
	Right *Tree
}

// Count the nodes in the given complete binary tree.
func (t *Tree) Count() int {
	// Only test the Left node (this binary tree is expected to be complete).
	if t.Left == nil {
		return 1
	}
	return 1 + t.Right.Count() + t.Left.Count()
}

// Create a complete binary tree of `depth` and return it as a pointer.
func NewTree(depth int) *Tree {
	if depth > 0 {
		return &Tree{Left: NewTree(depth - 1), Right: NewTree(depth - 1)}
	} else {
		return &Tree{}
	}
}
