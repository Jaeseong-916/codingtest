package main

import (
	"bufio"
	"fmt"
	"os"
)

func calc(n int) int {
	var rest int
	res := -1

	if n%5 == 0 {
		res = n / 5
	} else {
		for i := 0; i <= n/5; i++ {
			rest = n - (5 * i)
			if rest%3 == 0 {
				res = (rest / 3) + i
			}
		}
	}

	if res == -1 && n%3 == 0 {
		res = n / 3
	}

	return res
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	fmt.Fprintf(writer, "%d\n", calc(n))
}
