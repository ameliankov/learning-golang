package main

import (
	"fmt"
)

const target int = 9

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	var (
		left  int
		right int = len(nums) - 1
		sum   int
	)

	for left < right {
		sum = nums[left] + nums[right]
		fmt.Println(nums[left], "+", nums[right], "=", sum, "==", target)
		if target == sum {
			return []int{left, right}
		} else if target > sum {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}
