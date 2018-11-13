package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Print(sum(1, 2), "\n")
	result := sum(12, 2)
	fmt.Printf("%v\n", result)
	fmt.Printf("%v\n", mul(2,2))
}

func sum(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}
