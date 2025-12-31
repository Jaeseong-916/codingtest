package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var cnt = make([]int, n+1)

	for i := 2; i <= n; i++ {
		cnt[i] = cnt[i-1] + 1
		if i%2 == 0 {
			cnt[i] = min(cnt[i], cnt[i/2]+1)
		}
		if i%3 == 0 {
			cnt[i] = min(cnt[i], cnt[i/3]+1)
		}
	}
	fmt.Fprintf(writer, "%d\n", cnt[n])
}
