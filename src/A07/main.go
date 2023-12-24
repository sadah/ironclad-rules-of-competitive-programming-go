package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(D int64, N int64, L []int64, R []int64) {
	d := make([]int64, D+2)
	for i := int64(0); i < N; i++ {
		d[L[i]] += 1
		d[R[i]+1] -= 1
	}
	ans := int64(0)
	for i := int64(1); i <= D; i++ {
		ans += d[i]
		fmt.Println(ans)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var D int64
	scanner.Scan()
	D, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	L := make([]int64, N)
	R := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		L[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		R[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(D, N, L, R)
}
