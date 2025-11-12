package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n int

	fmt.Fscanln(reader, &n)
	fmt.Println("1")
	fmt.Println("0")
}
