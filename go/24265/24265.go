package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, res int

	fmt.Fscanln(reader, &n)
	for i := n - 1; i >= 0; i-- {
		res = res + i
	}
	fmt.Println(res)
	fmt.Println("2")
}
