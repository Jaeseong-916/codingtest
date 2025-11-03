package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var b string
	var n, res, conv int

	fmt.Fscanln(reader, &b, &n)

	for i := range b {
		if n < 10 {
			conv, _ = strconv.Atoi(string(b[i]))
			res = res + conv*int(math.Pow(float64(n), float64(len(b)-i-1)))
		} else {
			if int(b[i])-54 < 4 {
				conv, _ = strconv.Atoi(string(b[i]))
				res = res + conv*int(math.Pow(float64(n), float64(len(b)-i-1)))
			} else {
				conv = int(b[i] - 55)
				res = res + conv*int(math.Pow(float64(n), float64(len(b)-i-1)))
			}
		}
	}
	fmt.Println(res)
}
