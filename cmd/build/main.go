package main

import "github.com/gabivlj/candice/internals/build"

// Deploy: mv build candice && sudo mv candice /usr/local/bin
func main() {
	build.ExecuteProject()
}
