package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n int

	fmt.Fscanln(reader, &n)

	fmt.Println(int((math.Pow(2.0, float64(n)) + 1) * (math.Pow(2.0, float64(n)) + 1)))
}
