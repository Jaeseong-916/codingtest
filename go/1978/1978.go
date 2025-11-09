package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, res int
	fmt.Fscanln(reader, &n)

	for _ = range n {
		var input int
		fmt.Fscan(reader, &input)
		if IsPrime(input) {
			res++
		}
	}
	fmt.Println(res)
}

func IsPrime(val int) bool {
	if val <= 1 {
		return false
	}

	if val == 2 {
		return true
	}

	if val%2 == 0 {
		return false
	}

	lim := int(math.Sqrt(float64(val)))

	for i := 2; i <= lim; i++ {
		if val%i == 0 {
			return false
		}
	}
	return true
}
