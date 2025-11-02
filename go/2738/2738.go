package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, m int

	fmt.Fscanln(reader, &n, &m)

	a := make([][]int, n)
	b := make([][]int, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		str := strings.Fields(strings.TrimSuffix(line, "\n"))
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			a[i][j], _ = strconv.Atoi(str[j])
		}
	}
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		str := strings.Fields(strings.TrimSuffix(line, "\n"))
		b[i] = make([]int, m)
		for j := 0; j < m; j++ {
			b[i][j], _ = strconv.Atoi(str[j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", a[i][j]+b[i][j])
		}
		fmt.Println("")
	}
}
