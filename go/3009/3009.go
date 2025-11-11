package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var x1, x2, x3, x4, y1, y2, y3, y4 int

	fmt.Fscanln(reader, &x1, &y1)
	fmt.Fscanln(reader, &x2, &y2)
	fmt.Fscanln(reader, &x3, &y3)

	if x1 == x2 {
		x4 = x3
	} else if x2 == x3 {
		x4 = x1
	} else if x1 == x3 {
		x4 = x2
	}
	if y1 == y2 {
		y4 = y3
	} else if y2 == y3 {
		y4 = y1
	} else if y1 == y3 {
		y4 = y2
	}
	fmt.Printf("%d %d\n", x4, y4)
}
