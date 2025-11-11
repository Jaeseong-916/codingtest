package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b, c int
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)
	if a+b+c != 180 {
		fmt.Println("Error")
	} else {
		if a == 60 && b == 60 && c == 60 {
			fmt.Println("Equilateral")
		} else if a == b || a == c || b == c {
			fmt.Println("Isosceles")
		} else {
			fmt.Println("Scalene")
		}
	}
}
