package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b int

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	fmt.Println(a * b)
}
