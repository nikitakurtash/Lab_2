package main

func shellSort(nums []int) {
	n := len(nums)
	h := 1
	// Находим начальное значение h
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		// Производим вставку с инкрементом h
		for i := h; i < n; i++ {
			for j := i; j >= h && nums[j] < nums[j-h]; j -= h {
				nums[j], nums[j-h] = nums[j-h], nums[j]
			}
		}
		h /= 3
	}
}

func findKthLargest(nums []int, k int) int {
	shellSort(nums)
	return nums[len(nums)-k]
}
