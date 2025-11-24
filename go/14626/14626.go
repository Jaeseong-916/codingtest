package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input string
	var res int

	fmt.Fscanln(reader, &input)
	var sum, idx int
	for i := range len(input) {
		if input[i] == '*' {
			idx = i
			continue
		}
		if i%2 == 0 {
			num, _ := strconv.Atoi(string(input[i]))
			sum += num
		} else {
			num, _ := strconv.Atoi(string(input[i]))
			sum += num * 3
		}
	}
	if sum%10 == 0 {
		res = 0
	} else {
		if idx%2 == 0 {
			res = 10 - (sum % 10)
		} else {
			for i := 1; i < 10; i++ {
				if (3*i+sum)%10 == 0 {
					res = i
					break
				}
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", res)
}
