package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b, c, d, e int
	fmt.Fscanln(reader, &a, &b, &c, &d, &e)
	fmt.Println(((a * a) +
		(b * b) +
		(c * c) +
		(d * d) +
		(e * e)) % 10)
}
