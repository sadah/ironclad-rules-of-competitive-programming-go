package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(D int64, X int64, A []int64, Q int64, S []int64, T []int64) {

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
    var X int64
    scanner.Scan()
    X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    A := make([]int64, D-2+1)
    for i := int64(0); i < D-2+1; i++ {
        scanner.Scan()
        A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    var Q int64
    scanner.Scan()
    Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    S := make([]int64, Q)
    T := make([]int64, Q)
    for i := int64(0); i < Q; i++ {
        scanner.Scan()
        S[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        T[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(D, X, A, Q, S, T)
}
