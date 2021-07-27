package main

func main() {
	longestCommonSubsequence("ababab", "ab")
}

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	lcs := make([][]int, len(text1)+1)
	for i := range lcs {
		lcs[i] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = Max(lcs[i][j-1], lcs[i-1][j])
			}
		}
	}
	return lcs[len(text1)][len(text2)]
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
