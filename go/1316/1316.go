package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	var count, res int
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &count)

	for i := 0; i < count; i++ {
		fmt.Fscanln(reader, &input)

		cnt := make(map[uint8]bool)
		var prev uint8
		var isGroup = true

		for j := 0; j < len(input); j++ {
			_, exist := cnt[input[j]]
			if !exist {
				cnt[input[j]] = true
			} else if prev != input[j] {
				isGroup = false
			}
			prev = input[j]
		}
		if isGroup {
			res++
		}
	}
	fmt.Println(res)
}
