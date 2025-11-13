package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, cnt int

	fmt.Fscanln(reader, &n)
	var test_case = make([]string, n)
	var res = make([]int, n)

	for i := range n {
		fmt.Fscanln(reader, &test_case[i])
		for j := range len(test_case[i]) {
			if string(test_case[i][j]) == "O" {
				cnt++
				res[i] = res[i] + cnt
			} else if string(test_case[i][j]) == "X" {
				cnt = 0
			}
		}
		cnt = 0
	}

	for i := range n {
		fmt.Println(res[i])
	}
}
