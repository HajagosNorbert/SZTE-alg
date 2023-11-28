package main

import "fmt"

func oszd(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 2
	} else if n == 2 {
		return 4
	}

	return 2*oszd(n-1) + 1 + oszd(n-3)
}

func maxCost(cost [][]int, k, n int) int {
	if cost[0][0] == -1 {
		return -1
	}
	tc := make([][]int, k+1)
	for i := range tc {
		tc[i] = make([]int, n+1)
	}
	tc[0][0] = cost[0][0]

	for i := 1; i <= k; i++ {
		if cost[i][0] == -1 {
			for j := i; j <= k; j++ {
				tc[j][0] = -1
			}
			break
		}
		tc[i][0] = tc[i-1][0] + cost[i][0]
	}

	for j := 1; j <= n; j++ {
		if cost[0][j] == -1 {
			for i := j; i <= n; i++ {
				tc[0][i] = -1
			}
			break
		}
		tc[0][j] = tc[0][j-1] + cost[0][j]
	}

	for i := 1; i <= k; i++ {
		for j := 1; j <= n; j++ {
            if tc[i-1][j] == -1 && tc[i][j-1] == -1 {
                tc[i][j] = -1
            }
			if tc[i-1][j] == -1 {

                tc[i][j] = tc[i][j-1] + cost[i][j]
			} else if cost[i][j-1] == -1 {
                tc[i][j] = tc[i-1][j] + cost[i][j]
			}
			tc[i][j] = max(tc[i-1][j], tc[i][j-1]) + cost[i][j]
		}
	}
	return tc[k][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	cost := [][]int{{1, 1, -1}, {-1, 0, 1}, {1, -1, 1}}
	k, n := 2, 2
	result := maxCost(cost, k, n)
	fmt.Println(result)
}
