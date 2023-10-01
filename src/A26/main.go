package main

import (
	"bufio"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(Q int64, X []int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var Q int64
    scanner.Scan()
    Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    X := make([]int64, Q)
    for i := int64(0); i < Q; i++ {
        scanner.Scan()
        X[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(Q, X)
}
