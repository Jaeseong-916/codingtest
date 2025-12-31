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

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var passwd = make(map[string]string)
	for range n {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		inputStr := strings.Fields(line)

		passwd[inputStr[0]] = inputStr[1]
	}

	for range m {
		var tmp string
		fmt.Fscanln(reader, &tmp)
		fmt.Fprintf(writer, "%s\n", passwd[tmp])
	}
}
