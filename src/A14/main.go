package main

import (
	"bufio"
	"os"
	"strconv"
)

const YES = "Yes"

func solve(N int64, K int64, A []int64, B []int64, C []int64, D []int64) {

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
    var K int64
    scanner.Scan()
    K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    A := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    B := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    C := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        C[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    D := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        D[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, K, A, B, C, D)
}
