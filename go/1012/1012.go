package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(field [][]int, x, y, m, n int) {
	if x < 0 || x >= m || y < 0 || y >= n || field[x][y] != 1 {
		return
	}

	field[x][y] = 0
	dfs(field, x-1, y, m, n)
	dfs(field, x+1, y, m, n)
	dfs(field, x, y-1, m, n)
	dfs(field, x, y+1, m, n)
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	for range t {
		var n, m, k int
		fmt.Fscanln(reader, &m, &n, &k)

		// field 배열 초기화
		var field = make([][]int, m)
		for i := range field {
			field[i] = make([]int, n)
		}

		for range k {
			var x, y int
			fmt.Fscanln(reader, &x, &y)

			field[x][y] = 1
		}

		count := 0
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if field[i][j] == 1 {
					dfs(field, i, j, m, n)
					count++
				}
			}
		}
		fmt.Fprintf(writer, "%d\n", count)
	}
}
