package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, left, down, res int
	square := make([][]bool, 100)
	for i := range 100 {
		square[i] = make([]bool, 100)
	}

	fmt.Fscanln(reader, &n)

	for _ = range n {
		fmt.Fscanln(reader, &left, &down)
		// 색종이가 붙은 영역을 true로 표시
		for j := left; j < left+10; j++ {
			for k := down; k < down+10; k++ {
				square[j][k] = true
			}
		}
	}

	// true인 칸의 개수를 세서 넓이 계산
	for i := range 100 {
		for j := range 100 {
			if square[i][j] {
				res++
			}
		}
	}
	fmt.Println(res)
}
