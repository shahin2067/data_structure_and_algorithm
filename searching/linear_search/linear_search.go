package main

import "fmt"

func main() {

	nums := [...]int{1, 2, 21, 14, 27, 29, 11}
	var item int = 11
	var index = linearSearch(nums[:], item)
	if index != -1{
		fmt.Println(item, "Found at positions", index)
	} else {
		fmt.Println(item, "Not found")
	}
	
}

func linearSearch(nums []int, item int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == item {
			return i
		}
	}
	return -1 
}
