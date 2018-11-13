package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf("%v\n", sum(2,3))
	fmt.Printf("%v\n", sum(2,4))
	fmt.Printf("%v\n", mul(2,2))
	fmt.Printf("%v\n", mul(2,5))
	fmt.Printf("%v\n", mul(2,6))
}

func sum(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}
