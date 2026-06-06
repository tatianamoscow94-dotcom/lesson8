package main

import "fmt"

func main() {
	count := 0

	counter := func() int {
		count++
		return count
	}

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
