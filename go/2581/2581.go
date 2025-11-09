package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var m, n, min, sum int

	fmt.Fscanln(reader, &m)
	fmt.Fscanln(reader, &n)
	min = n

	for i := m; i <= n; i++ {
		if IsPrime(i) {
			if min > i {
				min = i
			}
			sum = sum + i
		}
	}
	if sum == 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(sum)
		fmt.Println(min)
	}
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
