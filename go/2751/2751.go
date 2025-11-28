package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int

	fmt.Fscanln(reader, &n)
	var input = make([]int, n)
	for i := range input {
		fmt.Fscanln(reader, &input[i])
	}

	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	for i := range input {
		fmt.Fprintf(writer, "%d\n", input[i])
	}
}
