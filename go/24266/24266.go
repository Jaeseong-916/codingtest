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
	fmt.Println(n * n * n)
	fmt.Println("3")
}
