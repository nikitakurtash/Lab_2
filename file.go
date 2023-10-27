package main

func partition(array []int, low, high int) int {
	i := (low - 1)
	pivot := array[high]
	for j := low; j < high; j++ {
		if array[j] <= pivot {
			i = i + 1
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[high] = array[high], array[i+1]
	return (i + 1)
}

func quickSort(array []int, low, high int) {
	if low < high {
		pi := partition(array, low, high)
		quickSort(array, low, pi-1)
		quickSort(array, pi+1, high)
	}
}

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}
