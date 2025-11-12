package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, res int

	fmt.Fscanln(reader, &n)

	for i := 0; i <= n-2; i++ {
		res = res + (n-2-i)*(i+1)
	}

	fmt.Println(res)
	fmt.Println("3")
}
