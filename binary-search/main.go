package main

import (
	"fmt"
)

const target int = 1

func main() {
	var nums = []int{-9, -1, 0, 7, 12, 13}
	pos := Search(nums, target)
	fmt.Println(pos)
}

func Search(nums []int, target int) int {
	var (
		low  int = 0
		mid  int = 0
		high int = len(nums) - 1
	)
	for low <= high {
		mid = (low + high) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
