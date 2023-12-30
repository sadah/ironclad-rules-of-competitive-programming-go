package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(H int64, W int64, X [][]int64, Q int64, A []int64, B []int64, C []int64, D []int64) {
	sums := make([][]int64, H+1)
	for i := range sums {
		sums[i] = make([]int64, W+1)
	}

	for i := int64(1); i <= H; i++ {
		sum := int64(0)
		for j := int64(1); j <= W; j++ {
			sum += X[i-1][j-1]
			sums[i][j] = sum
		}
	}

	for j := int64(1); j <= W; j++ {
		sum := int64(0)
		for i := int64(1); i <= H; i++ {
			sum += sums[i][j]
			sums[i][j] = sum
		}
	}

	for i := int64(0); i < Q; i++ {
		fmt.Println(sums[C[i]][D[i]] + sums[A[i]-1][B[i]-1] - sums[A[i]-1][D[i]] - sums[C[i]][B[i]-1])
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
	X := make([][]int64, H)
	for i := int64(0); i < H; i++ {
		X[i] = make([]int64, W)
	}
	for i := int64(0); i < H; i++ {
		for j := int64(0); j < W; j++ {
			scanner.Scan()
			X[i][j], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		}
	}
	var Q int64
	scanner.Scan()
	Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, Q)
	B := make([]int64, Q)
	C := make([]int64, Q)
	D := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		C[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		D[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(H, W, X, Q, A, B, C, D)
}
