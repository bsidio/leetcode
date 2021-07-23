package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, -4, 5, -5, 4, -1, -1, 2}

	fmt.Println(threesum(nums))
}

//sorted [-5 -4 -1 -1 1 2 4 5]

type List []int

func threesum(list []int) []List {
	var ansList []List
	lenth := len(list)
	sort.Ints(list)
	fmt.Println(list)
	for i := 0; i < lenth-2; i++ {

		low, high := i+1, lenth-1
		sum := 0 - list[i]
		for {
			if low >= high {
				break
			}

			switch {
			case list[low]+list[high] == sum:
				var ans List = []int{list[i], list[low], list[high]}
				fmt.Println(ans)
				ansList = append(ansList, ans)
				low++
				high = high - 1
			case list[low]+list[high] < sum:
				low++
			case list[low]+list[high] > sum:
				high = high - 1
			}

		}
	}
	return ansList
}

func threeSum2(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		for k, j := i+1, len(nums)-1; k < j; {
			n := nums[i] + nums[k] + nums[j]
			if n == 0 {
				result = append(result, []int{nums[i], nums[k], nums[j]})
				l := k
				for l < j && nums[l] == nums[k] {
					l++
				}
				k = l
			} else if n > 0 {
				j--
			} else {
				k++
			}
		}
	}
	return result
}
