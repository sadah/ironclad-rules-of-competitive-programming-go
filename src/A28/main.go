package main

import (
	"bufio"
	"os"
	"strconv"
)

const MOD = 10000

func solve(N int64, T []string, A []int64) {

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
    T := make([]string, N)
    A := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        T[i] = scanner.Text()
        scanner.Scan()
        A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, T, A)
}
