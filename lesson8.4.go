package main

import "fmt"

func sumAll(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	fmt.Println(sumAll(1, 2, 3))
	fmt.Println(sumAll(10, -2, 4, 7))
}
