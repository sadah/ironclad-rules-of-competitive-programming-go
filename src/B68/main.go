package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, M int64, P []int64, A []int64, B []int64) {

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
    var M int64
    scanner.Scan()
    M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    P := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        P[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    A := make([]int64, M)
    B := make([]int64, M)
    for i := int64(0); i < M; i++ {
        scanner.Scan()
        A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, M, P, A, B)
}
