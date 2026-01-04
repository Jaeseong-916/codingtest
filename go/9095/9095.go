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

	var dp = make([]int, 12)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4

	for i := 4; i < 12; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	for range t {
		var tmp int
		fmt.Fscanln(reader, &tmp)
		fmt.Fprintf(writer, "%d\n", dp[tmp])
	}
}
