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

	var n int
	fmt.Fscanln(reader, &n)

	var dp = make([]int, n+1)

	for i := range dp {
		dp[i] = 4
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			square := j * j
			if dp[i] > dp[i-square]+1 {
				dp[i] = dp[i-square] + 1
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", dp[n])
}
