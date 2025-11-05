package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, k, cnt int
	fmt.Fscanln(reader, &n, &k)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			cnt++
		}
		if cnt == k {
			fmt.Println(i)
			break
		}
	}
	if cnt < k {
		fmt.Println("0")
	}
}
