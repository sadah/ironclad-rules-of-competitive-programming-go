package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, S []int64, T []int64) {

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
    S := make([]int64, N)
    T := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        S[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        T[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, S, T)
}
