package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, m, max int

	fmt.Fscanln(reader, &n, &m)

	var input = make([]int, n)

	for i := range n {
		fmt.Fscan(reader, &input[i])
	}

	for i := range n - 2 {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if input[i]+input[j]+input[k] <= m {
					if input[i]+input[j]+input[k] >= max {
						max = input[i] + input[j] + input[k]
					}
				}
			}
		}
	}
	fmt.Println(max)
}
