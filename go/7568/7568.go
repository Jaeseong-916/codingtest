package main

import (
	"bufio"
	"fmt"
	"os"
)

func rank(person map[int][]int, cnt int) int {
	max_weight, max_hight := person[cnt][0], person[cnt][1]
	rk := 1
	for i := range person {
		if max_weight < person[i][0] && max_hight < person[i][1] {
			rk++
		}
	}
	return rk
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, x, y int
	fmt.Fscanln(reader, &n)
	person := make(map[int][]int, n)

	for i := range n {
		fmt.Fscanln(reader, &x, &y)
		person[i] = append(person[i], x)
		person[i] = append(person[i], y)
	}

	for i := range n {
		fmt.Fprintf(writer, "%d\n", rank(person, i))
	}
}
