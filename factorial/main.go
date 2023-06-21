package main

import "fmt"

const n int = 10

func main() {
	fmt.Println(factorial(n))
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
