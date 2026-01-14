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

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, "-")

	var res int
	if strings.Contains(parts[0], "+") {
		numparts := strings.Split(parts[0], "+")
		for _, val := range numparts {
			tmp, _ := strconv.Atoi(val)
			res += tmp
		}
	} else {
		tmp, _ := strconv.Atoi(parts[0])
		res += tmp
	}

	for i := 1; i < len(parts); i++ {
		sum := 0
		if strings.Contains(parts[i], "+") {
			numparts := strings.Split(parts[i], "+")
			for _, val := range numparts {
				tmp, _ := strconv.Atoi(val)
				sum += tmp
			}
		} else {
			tmp, _ := strconv.Atoi(parts[i])
			sum += tmp
		}
		res -= sum
	}
	fmt.Fprintf(writer, "%d\n", res)
}
