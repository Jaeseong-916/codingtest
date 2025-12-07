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

	var n int
	fmt.Fscanln(reader, &n)

	var queue = []int{}
	var res = []int{}

	for range n {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		parts := strings.Fields(line)

		var command = parts[0]

		if command == "push" && len(parts) == 2 {
			// push
			var val, _ = strconv.Atoi(parts[1])
			queue = append(queue, val)
		} else if command == "pop" {
			// pop
			if len(queue) == 0 {
				res = append(res, -1)
			} else {
				res = append(res, queue[0])
				queue = queue[1:]
			}
		} else if command == "size" {
			// size
			res = append(res, len(queue))
		} else if command == "empty" {
			// empty
			if len(queue) == 0 {
				res = append(res, 1)
			} else {
				res = append(res, 0)
			}
		} else if command == "front" {
			// front
			if len(queue) == 0 {
				res = append(res, -1)
			} else {
				res = append(res, queue[0])
			}
		} else if command == "back" {
			// back
			if len(queue) == 0 {
				res = append(res, -1)
			} else {
				res = append(res, queue[len(queue)-1])
			}
		}
	}

	for _, fmt_res := range res {
		fmt.Fprintf(writer, "%d\n", fmt_res)
	}
}
