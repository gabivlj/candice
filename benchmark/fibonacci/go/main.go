package main

import "fmt"

func fibonacci(x int) int64 {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func main() {
	for i := 0; i < 50; i++ {
		fmt.Println("value of", i, "is", fibonacci(i))
	}
}
