package main

import (
	"bufio"
	"fmt"
	"os"
)

func validate(board_validate []string) int {
	var cnt_start_w = 0
	var cnt_start_b = 0

	// W로 시작하는 패턴 체크
	for i := range board_validate {
		for j := range board_validate[i] {
			// (i+j)가 짝수면 첫 칸과 같은 색이어야 함
			if (i+j)%2 == 0 {
				if board_validate[i][j] != 'W' {
					cnt_start_w++
				}
			} else {
				if board_validate[i][j] != 'B' {
					cnt_start_w++
				}
			}
		}
	}

	// B로 시작하는 패턴 체크
	for i := range board_validate {
		for j := range board_validate[i] {
			// (i+j)가 짝수면 첫 칸과 같은 색이어야 함
			if (i+j)%2 == 0 {
				if board_validate[i][j] != 'B' {
					cnt_start_b++
				}
			} else {
				if board_validate[i][j] != 'W' {
					cnt_start_b++
				}
			}
		}
	}

	if cnt_start_b < cnt_start_w {
		return cnt_start_b
	} else {
		return cnt_start_w
	}
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var board_full = []string{}
	var res = 64

	for range n {
		var tmp string
		fmt.Fscanln(reader, &tmp)
		board_full = append(board_full, tmp)
	}

	// 모든 가능한 8x8 보드를 체크
	for i := 0; i <= n-8; i++ {
		for j := 0; j <= m-8; j++ {
			var board_validate = []string{}
			// 세로 8줄을 가져와서 각 줄에서 가로 8칸만 자르기
			for k := i; k < i+8; k++ {
				board_validate = append(board_validate, board_full[k][j:j+8])
			}
			tmp := validate(board_validate)
			if tmp < res {
				res = tmp
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", res)
}
