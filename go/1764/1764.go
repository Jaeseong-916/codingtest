package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var see_none = make(map[string]bool)
	for range n {
		tmp, _ := reader.ReadString('\n')
		tmp = strings.TrimSpace(tmp)
		see_none[tmp] = true
	}

	var res = []string{}
	for range m {
		tmp, _ := reader.ReadString('\n')
		tmp = strings.TrimSpace(tmp)
		if see_none[tmp] {
			res = append(res, tmp)
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	fmt.Fprintf(writer, "%d\n", len(res))
	for _, str := range res {
		fmt.Fprintf(writer, "%s\n", str)
	}
}
