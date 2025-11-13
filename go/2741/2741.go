package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n int

	fmt.Fscanln(reader, &n)
	for i := range n {
		fmt.Println(i + 1)
	}
}
