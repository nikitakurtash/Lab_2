package main

import (
	"fmt"
)

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]
	var less, greater []int
	for _, num := range nums[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}

func findKthLargest(nums []int, k int) int {
	sortedNums := quickSort(nums)
	return sortedNums[len(sortedNums)-k]
}

func main() {
	nums := []int{5, 2, 4, 1, 3}
	k := 3
	fmt.Println()
	fmt.Println(findKthLargest(nums, k))
}
