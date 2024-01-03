package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(A int64, B int64) {
	for i := A; i <= B; i++ {
		if 100%i == 0 {
			fmt.Println(YES)
			return
		}
	}
	fmt.Println(NO)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A int64
	scanner.Scan()
	A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(A, B)
}
