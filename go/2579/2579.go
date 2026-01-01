package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	stairs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &stairs[i])
	}

	if n == 1 {
		fmt.Fprintf(writer, "%d\n", stairs[1])
		return
	}
	if n == 2 {
		fmt.Fprintf(writer, "%d\n", stairs[1]+stairs[2])
		return
	}

	dp := make([]int, n+1)
	dp[1] = stairs[1]
	dp[2] = stairs[1] + stairs[2]

	for i := 3; i <= n; i++ {
		dp[i] = max(dp[i-3]+stairs[i-1]+stairs[i], dp[i-2]+stairs[i])
	}

	fmt.Fprintf(writer, "%d\n", dp[n])
}
