package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, X int64, A string) {

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
    var X int64
    scanner.Scan()
    X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var A string
    scanner.Scan()
    A = scanner.Text()
	solve(N, X, A)
}
