package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	var a, b, c, plus int

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)

	plus, _ = strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	fmt.Println(a + b - c)
	fmt.Println(plus - c)
}
