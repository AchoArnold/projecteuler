package projecteuler

func MaximumPathSumII(input [][]int) int {
	for i := len(input) - 2; i >= 0; i-- {
		for j := 0; j < len(input[i]); j++ {
			if input[i+1][j] > input[i+1][j+1] {
				input[i][j] += input[i+1][j]
			} else {
				input[i][j] += input[i+1][j+1]
			}
		}
	}
	return input[0][0]
}
