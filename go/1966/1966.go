package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var cnt_test int
	fmt.Fscanln(reader, &cnt_test)

	for range cnt_test {
		var n, m int
		fmt.Fscanln(reader, &n, &m)

		// 2차원 배열의 queue
		var queue = make([][]int, n)
		var input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		strNum := strings.Fields(input)

		for i, str := range strNum {
			tmp, err := strconv.Atoi(str)
			if err != nil {
				continue
			}
			queue[i] = []int{i, tmp}
		}

		// 1. 우선순위 체크
		// 2. 제일 클 때, queue[i][0] == m 이라면 cnt++, break. 아니라면 cnt++, out
		// 3. 아니라면 맨 뒤로 queue 이동

		var cnt int
		for len(queue) > 0 {
			// 최대 우선순위 확인
			var max = queue[0][1]
			for _, arr := range queue {
				if arr[1] > max {
					max = arr[1]
				}
			}
			// 제일 앞이 최대 우선순위인지 확인
			if queue[0][1] == max {
				cnt++
				if queue[0][0] == m {
					fmt.Fprintf(writer, "%d\n", cnt)
					break
				}
				// 출력 후 제거
				queue = queue[1:]
			} else {
				// 아니라면 맨 뒤로
				queue = append(queue, queue[0])
				queue = queue[1:]
			}
		}

	}
}
