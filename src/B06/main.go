package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A []int64, Q int64, L []int64, R []int64) {
	a := make([]int64, N+1)
	for i := int64(1); i <= N; i++ {
		a[i] = a[i-1] + A[i-1]
	}
	for i := int64(0); i < Q; i++ {
		cnt := R[i] - L[i] + 1
		ret := a[R[i]] - a[L[i]-1]
		if cnt == ret*2 {
			fmt.Println("draw")
		} else if cnt > ret*2 {
			fmt.Println("lose")
		} else {
			fmt.Println("win")
		}
	}
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
	var Q int64
	scanner.Scan()
	Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	L := make([]int64, Q)
	R := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		L[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		R[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A, Q, L, R)
}
