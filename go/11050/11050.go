package main

import (
	"bufio"
	"fmt"
	"os"
)

func facto(input int) int {
	output := 1
	for i := input; i > 0; i-- {
		output *= i
	}
	return output
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	res := facto(n) / (facto(k) * facto(n-k))

	fmt.Fprintf(writer, "%d\n", res)
}
