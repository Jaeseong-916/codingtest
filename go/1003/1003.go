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

	// 0과 1의 호출 횟수를 미리 계산 (최대 40까지)
	count0 := make([]int, 41)
	count1 := make([]int, 41)

	// 초기값 설정
	count0[0], count1[0] = 1, 0
	count0[1], count1[1] = 0, 1

	// 미리 계산
	for i := 2; i <= 40; i++ {
		count0[i] = count0[i-1] + count0[i-2]
		count1[i] = count1[i-1] + count1[i-2]
	}

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		fmt.Fprintf(writer, "%d %d\n", count0[n], count1[n])
	}
}
