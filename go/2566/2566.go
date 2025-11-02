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
	var max, row, col int = -1, 1, 1 // max를 -1로, row/col을 1로 초기화
	var st int

	for i := range 9 {
		line, _ := reader.ReadString('\n')
		str := strings.Fields(strings.TrimSpace(line))
		for j := range 9 {
			st, _ = strconv.Atoi(str[j])
			if st > max {
				max = st
				row = i + 1
				col = j + 1
			}
		}
	}
	fmt.Println(max)
	fmt.Printf("%d %d\n", row, col)
}

// var arr [][]int = make([][]int, 9)
// arr[i] = make([]int, 9)
// arr[i][j], _ = strconv.Atoi(str[j])

// for i := range 9 {
// 	for idx, j := range arr[i] {
// 		if j > max {
// 			max = j
// 			row = i + 1
// 			col = idx + 1
// 		}
// 	}
// }
