package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b, c int
	var res []string

	for {
		fmt.Fscanln(reader, &a, &b, &c)
		if a == 0 && b == 0 && c == 0 {
			break
		}
		if (a*a)+(b*b) == c*c {
			res = append(res, "right")
		} else if (a*a)+(c*c) == b*b {
			res = append(res, "right")
		} else if (b*b)+(c*c) == a*a {
			res = append(res, "right")
		} else {
			res = append(res, "wrong")
		}
	}
	for i := range res {
		fmt.Println(res[i])
	}
}
