package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strnum := strings.Fields(input)

	var time = []int{}
	for _, val := range strnum {
		tmp, _ := strconv.Atoi(val)
		time = append(time, tmp)
	}

	slices.Sort(time)

	var res int
	for i := len(time) - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			res += time[j]
		}
	}
	fmt.Fprintf(writer, "%d\n", res)
}
