package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, no int
	var n_str string

	fmt.Fscanln(reader, &n)

	for i := 1; i <= n; i++ {
		n_str = strconv.Itoa(i)
		var dassemble int = 0
		for j := range n_str {
			var n_tmp, _ = strconv.Atoi(string(n_str[j]))
			dassemble = dassemble + n_tmp
		}
		if dassemble+i == n {
			fmt.Println(i)
			no++
			break
		}
	}
	if no == 0 {
		fmt.Println(0)
	}

}
