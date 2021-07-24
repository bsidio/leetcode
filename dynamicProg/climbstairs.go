package main

import "fmt"

func main() {

	fmt.Println(ClimbStairs(5))
}

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}
