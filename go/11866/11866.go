package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	var arr = []int{}
	var res = []int{}

	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	index := 0
	for len(arr) > 0 {
		index = (index + k - 1) % len(arr)

		res = append(res, arr[index])
		arr = append(arr[:index], arr[index+1:]...)

		if len(arr) > 0 {
			index = index % len(arr)
		}
	}

	fmt.Fprintf(writer, "<")
	for i, v := range res {
		if i == len(res)-1 {
			fmt.Fprintf(writer, "%d", v)
		} else {
			fmt.Fprintf(writer, "%d, ", v)
		}
	}
	fmt.Fprintf(writer, ">\n")
}
