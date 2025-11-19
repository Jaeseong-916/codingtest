package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var input string
	var res []string

	for {
		fmt.Fscanln(reader, &input)
		if input == "0" {
			break
		}
		if len(input) == 1 {
			res = append(res, "yes")
			continue
		}
		var check = 0
		for i := range len(input) / 2 {
			if input[i] != input[len(input)-1-i] {
				res = append(res, "no")
				check++
				break
			}
		}
		if check != 1 {
			res = append(res, "yes")
		}
	}

	for i := range res {
		fmt.Println(res[i])
	}
}
