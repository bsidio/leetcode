package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 5}

	containsDuplicate2(nums)
}

//First method Passed
func containsDuplicate(nums []int) bool {
	var l = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		l[nums[i]] = i
	}
	if len(nums) > len(l) {
		fmt.Println("true")

		return true
	}
	return false
}

//Second Method

func containsDuplicate2(nums []int) bool {
	var l = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i, ok := l[nums[i]]; ok {
			fmt.Println(i)
			fmt.Println("true")

			return true
		}
		l[nums[i]] = i
	}
	fmt.Println("false")

	return false

}
