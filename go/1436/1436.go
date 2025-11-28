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

	var n, res int
	var cnt = 1

	fmt.Fscanln(reader, &n)

	for i := 1; cnt != n+1; i++ {
		if strings.Contains(strconv.Itoa(i), "666") {
			res = i
			cnt++
		}
	}

	fmt.Fprintf(writer, "%d\n", res)
}
