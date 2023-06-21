package main

import "fmt"

const n int = 5

func main() {
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
