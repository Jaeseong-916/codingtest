package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	var res []int
	for range t {
		var n int
		var clothes = make(map[string]int)
		fmt.Fscanln(reader, &n)
		for range n {
			var tmp, _ = reader.ReadString('\n')
			tmp = strings.TrimSpace(tmp)
			parts := strings.Fields(tmp)
			clothes[parts[1]]++
		}
		var cnt = 1
		for _, val := range clothes {
			cnt = cnt * (val + 1)
		}
		res = append(res, cnt-1)
	}
	for _, val := range res {
		fmt.Fprintf(writer, "%d\n", val)
	}
}
