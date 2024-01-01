package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(H int64, W int64, N int64, A []int64, B []int64, C []int64, D []int64) {
	board := make([][]int64, H+1)
	ans := make([][]int64, H+1)
	for i := range ans {
		board[i] = make([]int64, W+1)
		ans[i] = make([]int64, W+1)
	}

	for i := int64(0); i < N; i++ {
		board[A[i]][B[i]] += 1
		board[A[i]][D[i]+1] -= 1
		board[C[i]+1][B[i]] -= 1
		board[C[i]+1][D[i]+1] += 1
	}

	for i := int64(1); i <= H; i++ {
		for j := int64(1); j <= W; j++ {
			ans[i][j] = ans[i-1][j] + ans[i][j-1] - ans[i-1][j-1] + board[i][j]
		}
	}

	for i := int64(1); i <= H; i++ {
		for j := int64(1); j <= W; j++ {
			if j != W {
				fmt.Print(ans[i][j], " ")
			} else {
				fmt.Print(ans[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var H int64
	scanner.Scan()
	H, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var W int64
	scanner.Scan()
	W, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	B := make([]int64, N)
	C := make([]int64, N)
	D := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		C[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		D[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(H, W, N, A, B, C, D)
}
