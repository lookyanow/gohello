package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf("This is golang test\n")
	fmt.Print(sum(1, 2), "\n")
	result := sum(3, 2)
	fmt.Printf("%v\n", result)
}

func sum(a int, b int) int {
	return a + b
}
