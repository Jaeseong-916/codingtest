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

	var n, k, res int
	fmt.Fscanln(reader, &n, &k)

	var coin = []int{}

	for range n {
		var tmp int
		fmt.Fscanln(reader, &tmp)
		coin = append(coin, tmp)
	}

	for i := len(coin) - 1; i >= 0; i-- {
		if k == 0 {
			break
		} else if k < coin[i] {
			continue
		} else {
			if k%coin[i] != 0 && k/coin[i] != 0 {
				res += k / coin[i]
				k -= coin[i] * (k / coin[i])
			} else if k%coin[i] == 0 && k/coin[i] != 0 {
				res += k / coin[i]
				k -= coin[i] * (k / coin[i])
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", res)
}
