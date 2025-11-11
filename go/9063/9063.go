package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, x_min, y_min, x_max, y_max, x, y int
	fmt.Fscanln(reader, &n)

	x_max = -100000
	y_max = -100000
	x_min = 100000
	y_min = 100000

	for range n {
		fmt.Fscanln(reader, &x, &y)
		if x_max <= x {
			x_max = x
		}
		if x_min >= x {
			x_min = x
		}
		if y_max <= y {
			y_max = y
		}
		if y_min >= y {
			y_min = y
		}
	}
	fmt.Println((x_max - x_min) * (y_max - y_min))
}
