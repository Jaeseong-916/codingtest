package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a1, a0, c, n0 int
	fmt.Fscanln(reader, &a1, &a0)
	fmt.Fscanln(reader, &c)
	fmt.Fscanln(reader, &n0)

	if (a1*n0)+a0 <= c*n0 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
