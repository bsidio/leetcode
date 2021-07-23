package main

import "fmt"

func main() {
	nums := []int{4, 2, 3, 5}

	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	l := len(nums)
	gl := make([]int, l)
	gl[0] = 1
	for i := 1; i < l; i++ {
		gl[i] = gl[i-1] * nums[i-1]
	}
	gr := nums[l-1]
	for i := l - 2; i >= 0; i-- {
		gl[i] *= gr
		gr *= nums[i]
	}

	return gl
}
func productExceptSelf2(nums []int) []int {
	n := len(nums)
	var gl = make([]int, len(nums))
	//	var gr =make([]int,len(nums))
	fmt.Println(nums)
	gl[0] = 1
	for i := 1; i < len(nums); i++ {
		gl[i] = nums[i-1] * gl[i-1]
		//fmt.Println(nums[i-1], gl[i-1], gl[i])
	}
	// fmt.Println(gl)
	gr := nums[n-1]
	// fmt.Println(gr)
	fmt.Println(gl)
	for i := n - 2; i >= 0; i-- {

		fmt.Println(gl[i], gr)
		fmt.Println(nums)
		fmt.Println(gl)

		gl[i] *= gr
		fmt.Println(gl)

		gr *= nums[i]
		// fmt.Println(gr, nums[i])

	}

	return gl
}
