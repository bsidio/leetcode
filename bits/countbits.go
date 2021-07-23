package main

import "fmt"

func main() {

	fmt.Println(countBits(5))

}

func countBits(n int) []int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = f[i>>1] + (i & 1)
	}
	return f
}
