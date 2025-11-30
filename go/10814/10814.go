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

	var age int
	var name string
	var table = make(map[int][]string)
	for range n {
		fmt.Fscanln(reader, &age, &name)
		table[age] = append(table[age], name)
	}

	for i := 1; i <= 200; i++ {
		if table[i] != nil {
			for j := range table[i] {
				fmt.Fprintf(writer, "%d %s\n", i, table[i][j])
			}
		}
	}
}
