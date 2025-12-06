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

	var k, res int
	fmt.Fscanln(reader, &k)

	var stack = []int{}

	for range k {
		var input int
		fmt.Fscanln(reader, &input)
		if input != 0 {
			stack = append(stack, input)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	for i := range stack {
		res += stack[i]
	}
	fmt.Fprintf(writer, "%d\n", res)
}
