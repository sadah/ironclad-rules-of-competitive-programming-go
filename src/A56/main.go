package main

import (
	"bufio"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(N int64, Q int64, S string, a []int64, b []int64, c []int64, d []int64) {

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
    var Q int64
    scanner.Scan()
    Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var S string
    scanner.Scan()
    S = scanner.Text()
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
	solve(N, Q, S, a, b, c, d)
}
