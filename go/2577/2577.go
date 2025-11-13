package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b, c int
	var mul_res string
	res := make(map[string]int)

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)

	mul_res = strconv.Itoa((a * b * c))

	for i := range mul_res {
		res[string(mul_res[i])]++
	}

	for i := range 10 {
		if res[strconv.Itoa(i)] != 0 {
			fmt.Println(res[strconv.Itoa(i)])
		} else {
			fmt.Println("0")
		}
	}
}
