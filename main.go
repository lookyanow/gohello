package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf("This is golang test\n")
	fmt.Print(sum(1, 2), "\n")
	result := sum(3, 2)
	fmt.Printf("%v\n", result)
	fmt.Printf("%d\n", mul(2,2))
	var res bool = false
	var str string

	str = os.Args[1]
	res = str == "close" || str == "merge"
	fmt.Printf("%v\n", res)

}

func sum(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}