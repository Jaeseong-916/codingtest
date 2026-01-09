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

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	num_str := strings.Fields(line)
	var num []int
	for _, val := range num_str {
		tmp, _ := strconv.Atoi(val)
		num = append(num, tmp)
	}

	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + num[i-1]
	}

	var res []int
	for range m {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		num_str = strings.Fields(line)

		var add []int
		for _, val := range num_str {
			tmp, _ := strconv.Atoi(val)
			add = append(add, tmp)
		}
		sum := prefix[add[1]] - prefix[add[0]-1]
		res = append(res, sum)
	}

	for _, val := range res {
		fmt.Fprintf(writer, "%d\n", val)
	}
}
