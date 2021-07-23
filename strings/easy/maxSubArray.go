package main

import "fmt"

func main() {
	t := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(t))

}
func maxSubArray(nums []int) int {
	me := nums[0]
	mt := nums[0]

	for _, num := range nums[1:] {
		fmt.Println("num", num)
		fmt.Println("me", me)
		fmt.Println("mt", mt)
		me = max(me+num, num)
		mt = max(me, mt)
	}

	return mt
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
