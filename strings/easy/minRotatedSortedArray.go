package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5, 1}
	//nums2 := []int{4, 1, 2, 3}

	findMin(nums)
}
func findMin(nums []int) int {

	low, high := 0, len(nums)-1
	for low < high {
		if nums[low] < nums[high] {
			return nums[low]
		}
		mid := low + (high-low)>>1
		fmt.Println("mid", mid)
		if nums[mid] >= nums[low] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}
