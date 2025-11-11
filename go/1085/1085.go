package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var x, y, w, h int

	fmt.Fscanln(reader, &x, &y, &w, &h)

	if w/2 >= x && h/2 <= y { // 1사분면
		if x < h-y {
			fmt.Println(x)
		} else {
			fmt.Println(h - y)
		}
	} else if w/2 < x && h/2 < y { // 2사분면
		if w-x < h-y {
			fmt.Println(w - x)
		} else {
			fmt.Println(h - y)
		}
	} else if w/2 <= x && h/2 >= y { // 3사분면
		if w-x < y {
			fmt.Println(w - x)
		} else {
			fmt.Println(y)
		}
	} else if w/2 > x && h/2 > y { // 4사분면
		if x < y {
			fmt.Println(x)
		} else {
			fmt.Println(y)
		}
	} else {
		if x < y {
			fmt.Println(x)
		} else {
			fmt.Println(y)
		}
	}
}
