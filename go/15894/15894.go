package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, plus, minus int
	fmt.Fscanln(reader, &n)
	for i := range n {
		plus = plus + i + 1
		minus = minus + i
	}
	fmt.Println(4*plus - (2 * 2 * minus))
}
