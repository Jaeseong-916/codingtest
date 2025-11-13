package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var t, h, w, n int

	fmt.Fscanln(reader, &t)
	var res = make([]int, t)

	for i := range t {
		fmt.Fscanln(reader, &h, &w, &n)

	}
}

// 처음부터 h만큼 1씩 더하면서 진행.
// h 이후 w-h만큼 유지하면서 진행.
// h + w - h 이후 h만큼 1씩 감소하면서 진행.
