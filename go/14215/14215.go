package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b, c int
	fmt.Fscanln(reader, &a, &b, &c)
	if a+b <= c {
		c = a + b - 1
	} else if a+c <= b {
		b = a + c - 1
	} else if b+c <= a {
		a = b + c - 1
	}
	fmt.Println(a + b + c)
}
