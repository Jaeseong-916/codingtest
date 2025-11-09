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
	if n == 1 {
		fmt.Println("")
	} else {
		Prime_Factorization(n)
	}
}

func Prime_Factorization(val int) {
	for i := 2; i <= val; i++ {
		if val%i == 0 {
			fmt.Println(i)
			Prime_Factorization(val / i)
			break
		}
	}
}
