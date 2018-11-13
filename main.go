package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf("%v\n", sum(2,3))
	fmt.Printf("%v\n", mul(2,2))
}

func sum(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}
