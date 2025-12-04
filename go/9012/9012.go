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

	for range n {
		line, _ := reader.ReadString('\n')

		stack := []rune{}
		vps := true

		for _, ch := range line {
			if ch == '(' {
				stack = append(stack, ch)
			} else if ch == ')' {
				if len(stack) == 0 || stack[len(stack)-1] != '(' {
					vps = false
					break
				}
				stack = stack[:len(stack)-1]
			}
		}
		if len(stack) == 0 && vps {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}
}
