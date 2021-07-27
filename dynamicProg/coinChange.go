package main

func main() {

}
func coinChange(coins []int, amount int) int {
	//  sort.Ints(coins)
	// lent:=len(coins)
	amt := make([]int, amount+1)
	// amt[0]=0
	for i := 0; i < len(amt); i++ {
		amt[i] = amount + 1
	}
	amt[0] = 0

	// fmt.Println(amt[0])
	// fmt.Println(int(math.Min(float64(amt[0]),float64(amt[1]))))
	for _, coin := range coins {
		for i := 1; i <= amount; i++ {

			// fmt.Println(i,i-coin)
			if i-coin >= 0 {
				//fmt.Println("s",amt[i],amt[i-coin],i-coin)
				amt[i] = Getmin(amt[i], amt[i-coin]+1)
				// fmt.Println(amt[i],amt[i-coin]+1,Getmin(amt[i],amt[i-coin]+1))
			}

		}

	}
	if amt[amount] > amount {
		return -1
	}
	// fmt.Println(amt)
	return amt[amount]
}

// sort.

func Getmin(a int, b int) int {
	if a < b {
		//fmt.Println("a",a)
		return a
	} else {
		// fmt.Println("b",b)
		return b
	}
}
