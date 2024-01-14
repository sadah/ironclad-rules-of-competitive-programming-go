package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, X []int64, Y []int64, Q int64, a []int64, b []int64, c []int64, d []int64) {
	s, t := make([][]int64, 1501), make([][]int64, 1501)
	for i := range s {
		s[i] = make([]int64, 1501)
		t[i] = make([]int64, 1501)
	}

	for i := int64(0); i < N; i++ {
		s[X[i]][Y[i]]++
	}

	for i := int64(1); i <= 1500; i++ {
		for j := int64(1); j <= 1500; j++ {
			t[i][j] = t[i][j-1] + s[i][j]
		}
	}

	for i := int64(1); i <= 1500; i++ {
		for j := int64(1); j <= 1500; j++ {
			t[i][j] = t[i-1][j] + t[i][j]
		}
	}

	for i := int64(0); i < Q; i++ {
		fmt.Println(t[c[i]][d[i]] + t[a[i]-1][b[i]-1] - t[a[i]-1][d[i]] - t[c[i]][b[i]-1])
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
	X := make([]int64, N)
	Y := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		X[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		Y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	var Q int64
	scanner.Scan()
	Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	a := make([]int64, Q)
	b := make([]int64, Q)
	c := make([]int64, Q)
	d := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		c[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		d[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, X, Y, Q, a, b, c, d)
}
