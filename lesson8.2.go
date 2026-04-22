package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	x, y := 15, 4

	resAdd := applyOperation(x, y, add)
	resSub := applyOperation(x, y, sub)
	resMul := applyOperation(x, y, mul)

	fmt.Printf("%d + %d = %d\n", x, y, resAdd)
	fmt.Printf("%d - %d = %d\n", x, y, resSub)
	fmt.Printf("%d * %d = %d\n", x, y, resMul)
}
