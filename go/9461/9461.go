package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	var dp = make([]int, 101)

	dp[1] = 1
	dp[2] = 1
	dp[3] = 1

	for i := 4; i <= 100; i++ {
		dp[i] = dp[i-2] + dp[i-3]
	}

	for range t {
		var n int
		fmt.Fscanln(reader, &n)
		fmt.Fprintf(writer, "%d\n", dp[n])
	}
}
