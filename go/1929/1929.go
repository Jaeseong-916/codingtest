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

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var validate = make([]bool, 1000001)
	validate[0] = true
	validate[1] = true

	for i := 2; i < 1000000; i++ {
		if validate[i] {
			continue
		}
		for j := i * i; j < 1000000; j += i {
			if validate[j] {
				continue
			}
			if j%i == 0 {
				validate[j] = true
			}
		}
	}
	for i := n; i <= m; i++ {
		if validate[i] {
			continue
		}
		fmt.Fprintf(writer, "%d\n", i)
	}
}
