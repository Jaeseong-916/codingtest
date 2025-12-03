package main

import (
	"bufio"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		// 한 줄 입력 받기
		line, _ := reader.ReadString('\n')

		// 종료 조건: 온점 하나만 있는 경우
		if line[0] == '.' && (len(line) == 2 || len(line) == 1) {
			break
		}

		// 스택으로 사용할 슬라이스
		stack := []rune{}
		balanced := true

		// 문자열을 한 글자씩 검사
		for _, ch := range line {
			if ch == '(' || ch == '[' {
				// 여는 괄호는 스택에 push
				stack = append(stack, ch)
			} else if ch == ')' {
				// 닫는 소괄호: 스택이 비어있거나 top이 '('가 아니면 불균형
				if len(stack) == 0 || stack[len(stack)-1] != '(' {
					balanced = false
					break
				}
				// 짝이 맞으면 pop
				stack = stack[:len(stack)-1]
			} else if ch == ']' {
				// 닫는 대괄호: 스택이 비어있거나 top이 '['가 아니면 불균형
				if len(stack) == 0 || stack[len(stack)-1] != '[' {
					balanced = false
					break
				}
				// 짝이 맞으면 pop
				stack = stack[:len(stack)-1]
			}
			// 그 외 문자(알파벳, 공백, 마침표)는 무시
		}

		// 모든 괄호를 검사한 후 스택이 비어있고 balanced가 true면 균형잡힘
		if len(stack) == 0 && balanced {
			writer.WriteString("yes\n")
		} else {
			writer.WriteString("no\n")
		}
	}
}
