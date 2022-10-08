package main

import "fmt"

func main() {

	nums := [...]int{1, 2, 21, 14, 27, 29, 31}
	var x int = 31
	var n int = len(nums)
	var min, max = 0, n - 1

	var index = binarySearch(nums[:], min, max, x)

	if index != -1 {
		fmt.Println(x, "is Found at positions", index)
	} else {
		fmt.Println(x, "is not found")
	}
}

func binarySearch(nums []int, left int, right int, item int) int {
	if left <= right {
		var middle int = (left + (right-left)/2)

		if nums[middle] == item {
			return middle
		}

		if nums[middle] < item {
			return binarySearch(nums, middle+1, right, item)
		}

		return binarySearch(nums, left, middle-1, item)
	}

	return -1
}
