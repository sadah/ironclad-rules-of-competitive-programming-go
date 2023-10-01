package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(T int64, P []int64, Q []int64, R []int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var T int64
    scanner.Scan()
    T, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    P := make([]int64, T)
    Q := make([]int64, T)
    R := make([]int64, T)
    for i := int64(0); i < T; i++ {
        scanner.Scan()
        P[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        Q[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        R[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(T, P, Q, R)
}
