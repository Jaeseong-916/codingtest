package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var a int
	fmt.Fscanln(reader, &a)

	for i := 0; i < a-1; i++ {
		for j := a - 1; j > i; j-- {
			fmt.Fprint(writer, " ")
		}
		for j := 0; j < i; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprint(writer, "*")
		for j := 0; j < i; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer, "")
	}
	for i := 0; i < 2*a-1; i++ {
		fmt.Fprint(writer, "*")
	}
	fmt.Fprintln(writer, "")
	for i := 0; i < a-1; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Fprint(writer, " ")
		}
		for j := a - 2; j > i; j-- {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprint(writer, "*")
		for j := a - 2; j > i; j-- {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer, "")
	}
}
