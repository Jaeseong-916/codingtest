package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, b int

	fmt.Fscanln(reader, &n, &b)

	result := convertToBase(n, b)
	fmt.Println(result)
}

func convertToBase(n, b int) string {
	if n == 0 {
		return "0"
	}

	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := ""

	for n > 0 {
		remainder := n % b
		result = string(digits[remainder]) + result
		n = n / b
	}

	return result
}
