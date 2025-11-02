package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var str []string = make([]string, 5)

	for i := range 5 {
		str[i], _ = reader.ReadString('\n')
		str[i] = strings.TrimSuffix(str[i], "\n")
	}

	for i := range 15 {
		for j := range 5 {
			if len(str[j]) < i+1 {
				continue
			} else {
				fmt.Print(string(str[j][i]))
			}
		}
	}
}
