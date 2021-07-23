package main

import "fmt"

func main() {
	nums := []int{1, -4, 5, -5, 4}
	target := 0

	fmt.Println(twoSums(nums, target))
}

func twoSums(nums []int, target int) []int {
	//create a map to store the element of nums as key and their index as value
	result := make(map[int]int)
	//result["val"] :=make(map[int]int)
	//loop through the array nums
	for idx, v := range nums {
		//sub target-v to get remainder

		comp := target - v
		//search of remainder in array okay=true assign value of map to item(which here is the index of remainder from nums)
		if item, ok := result[comp]; ok {

			//if remainder in array return and array with
			return []int{item, idx}
		}
		//storing nums element as key and index as value
		result[v] = idx
	}
	fmt.Println(result)

	//nothing then return
	return []int{-1, 1}
}
