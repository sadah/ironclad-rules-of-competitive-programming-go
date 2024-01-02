package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A []int64, D int64, L []int64, R []int64) {
	p, q := make([]int64, N), make([]int64, N)

	p[0] = A[0]
	for i := int64(1); i < N; i++ {
		p[i] = max(p[i-1], A[i])
	}

	q[N-1] = A[N-1]
	for i := int64(N - 2); i >= 0; i-- {
		q[i] = max(q[i+1], A[i])
	}

	for i := int64(0); i < D; i++ {
		fmt.Println(max(p[L[i]-2], q[R[i]]))
	}
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	var D int64
	scanner.Scan()
	D, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	L := make([]int64, D)
	R := make([]int64, D)
	for i := int64(0); i < D; i++ {
		scanner.Scan()
		L[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		R[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A, D, L, R)
}
