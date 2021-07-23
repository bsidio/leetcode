package main

import (
	"fmt"
	"math"
)

func main() {
	t := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxProduct(t))

}
func maxProduct(nums []int) int {
	minnv := nums[0]
	maxv := nums[0]
	maxprod := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			minnv, maxv = maxv, minnv
		}
		minnv = int(math.Min(float64(nums[i]), float64(minnv*nums[i])))
		maxv = int(math.Max(float64(nums[i]), float64(maxv*nums[i])))
		if maxv > maxprod {
			maxprod = maxv
		}

	}
	return maxprod
}

func maxProduct2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	maxProd := 0
	currProd := 1

	for i := 0; i < n; i++ {
		currProd = currProd * nums[i]
		maxProd = max(maxProd, currProd)
		if currProd == 0 {
			currProd = 1
		}
	}

	currProd = 1
	for i := n - 1; i > 0; i-- {
		currProd = currProd * nums[i]
		maxProd = max(maxProd, currProd)
		if currProd == 0 {
			currProd = 1
		}
	}

	return maxProd
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
