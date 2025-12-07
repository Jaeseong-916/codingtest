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

	var n int
	var stack = []int{}
	var res = []int{}
	fmt.Fscanln(reader, &n)

	for range n {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		// 공백으로 분리
		parts := strings.Fields(line)

		if len(parts) > 0 {
			command := parts[0]

			if command == "push" && len(parts) == 2 {
				// push 명령어인 경우, 두 번째 요소가 숫자
				val, _ := strconv.Atoi(parts[1])
				stack = append(stack, val)
			} else if command == "top" {
				if len(stack) == 0 {
					res = append(res, -1)
				} else {
					res = append(res, stack[len(stack)-1])
				}
			} else if command == "pop" {
				if len(stack) == 0 {
					res = append(res, -1)
				} else {
					res = append(res, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
			} else if command == "size" {
				res = append(res, len(stack))
			} else if command == "empty" {
				if len(stack) == 0 {
					res = append(res, 1)
				} else {
					res = append(res, 0)
				}
			}
		}
	}
	for i := range res {
		fmt.Fprintf(writer, "%d\n", res[i])
	}
}
