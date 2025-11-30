package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var x, y int
	var coor = make([][]int, n)

	for i := range n {
		fmt.Fscanln(reader, &x, &y)
		coor[i] = append(coor[i], x, y)
	}

	sort.Slice(coor, func(i, j int) bool {
		if coor[i][1] == coor[j][1] {
			return coor[i][0] < coor[j][0]
		}
		return coor[i][1] < coor[j][1]
	})

	for i := range n {
		fmt.Fprintf(writer, "%d %d\n", coor[i][0], coor[i][1])
	}
}
