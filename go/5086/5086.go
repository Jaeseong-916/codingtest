package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var res = make([]string, 0)
	for {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		if a == 0 && b == 0 {
			break
		}
		if a < b && b%a == 0 {
			res = append(res, "factor")
		} else if a > b && a%b == 0 {
			res = append(res, "multiple")
		} else {
			res = append(res, "neither")
		}
	}
	for i := range res {
		fmt.Println(res[i])
	}
}
