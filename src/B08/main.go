package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, X []int64, Y []int64, Q int64, a []int64, b []int64, c []int64, d []int64) {

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
