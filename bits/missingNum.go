package main

func main() {

	//fmt.Println(missingNumber(5))

}

func missingNumber(nums []int) int {
	sum := 0
	total := 0
	for k, v := range nums {
		sum = sum + v
		total = total + k + 1
	}
	return total - sum
}
