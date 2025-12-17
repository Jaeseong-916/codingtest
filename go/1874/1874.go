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

	var input = []int{}
	for range n {
		var num int
		fmt.Fscanln(reader, &num)
		input = append(input, num)
	}

	var stack = []int{}
	var result = []string{}
	current := 1 // push할 다음 숫자

	for _, target := range input {
		// target까지 push
		for current <= target {
			stack = append(stack, current)
			result = append(result, "+")
			current++
		}

		// 스택의 top이 target과 같으면 pop
		if len(stack) > 0 && stack[len(stack)-1] == target {
			stack = stack[:len(stack)-1]
			result = append(result, "-")
		} else {
			// 불가능한 수열
			fmt.Fprintf(writer, "NO\n")
			return
		}
	}

	// 성공한 경우 결과 출력
	for _, op := range result {
		fmt.Fprintf(writer, "%s\n", op)
	}
}
