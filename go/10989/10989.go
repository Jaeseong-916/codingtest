package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int

	fmt.Fscanln(reader, &n)

	cnt := make([]int, 10001)

	for range n {
		var num int
		fmt.Fscanln(reader, &num)
		cnt[num]++
	}

	for i := 1; i < 10001; i++ {
		if cnt[i] > 0 {
			for range cnt[i] {
				fmt.Fprintln(writer, i)
			}
		}
	}
}
