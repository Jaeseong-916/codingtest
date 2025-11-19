package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b int
	var max_mod, min_mul int

	fmt.Fscanln(reader, &a, &b)

	min_mul = a * b

	if a > b {
		for i := 1; i <= b; i++ {
			if a%i == 0 && b%i == 0 {
				if i >= max_mod {
					max_mod = i
				}
			}
		}
		for i := a * b; i >= a; i-- {
			if i%a == 0 && i%b == 0 {
				if i <= min_mul {
					min_mul = i
				}
			}
		}
	} else if a < b {
		for i := 1; i <= a; i++ {
			if a%i == 0 && b%i == 0 {
				if i >= max_mod {
					max_mod = i
				}
			}
		}
		for i := a * b; i >= b; i-- {
			if i%a == 0 && i%b == 0 {
				if i <= min_mul {
					min_mul = i
				}
			}
		}
	} else {
		max_mod = a
		min_mul = a
	}
	fmt.Println(max_mod)
	fmt.Println(min_mul)
}
