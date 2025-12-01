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

	var n int
	fmt.Fscanln(reader, &n)

	var queue []int
	for i := 1; i <= n; i++ {
		queue = append(queue, i)
	}

	for len(queue) > 1 {
		// 맨 위 카드 버리기
		queue = queue[1:]
		// 그 다음 카드를 맨 아래로
		queue = append(queue, queue[0])
		queue = queue[1:]
	}

	fmt.Fprintf(writer, "%d\n", queue[0])
}
