package main

import "fmt"

/**
	Написать функцию, которая принимает массив чисел.
	Необходимо определить монотонный он или нет.
	Примеры:
	[1, 2, 3, 6] - true
	[6, 3, 2, 1] - true
	[5, 5] - true
	[1, 2, 5, 5, 5, 8, 9] - true
	[1, 2, 5, 5, 5, 2, 1] - false
	[1, 10, 6] - false
**/

func main() {
	var arr = []int{1, 10, 6}
	result := mono(arr)
	fmt.Println("Result:", result)
}

func mono(arr []int) bool {
	var asc, desc bool
	for i := 0; i <= len(arr)-2; i++ {
		if arr[i] <= arr[i+1] {
			if desc {
				return asc
			} else {
				asc = true
				continue
			}
		}
		if arr[i] >= arr[i+1] {
			if asc {
				return desc
			} else {
				desc = true
				continue
			}
		}
	}
	return true
}
