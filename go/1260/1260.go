package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)

	var n, m, v int
	fmt.Fscanln(reader, &n, &m, &v)

	var graph = make([][]int, n+1)
	for range m {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		numstr := strings.Fields(line)

		var numarr []int
		for _, val := range numstr {
			tmp, _ := strconv.Atoi(val)
			numarr = append(numarr, tmp)
		}
		graph[numarr[0]] = append(graph[numarr[0]], numarr[1])
		graph[numarr[1]] = append(graph[numarr[1]], numarr[0])
	}
	for i := 1; i <= n; i++ {
		sort.Ints(graph[i])
	}

	visited := make([]bool, n+1)
	dfs(v, graph, visited)

	fmt.Println()

	visited = make([]bool, n+1)
	bfs(v, graph, visited)

	fmt.Println()
}

func dfs(v int, graph [][]int, visited []bool) {
	visited[v] = true

	fmt.Print(v, " ")

	for _, next := range graph[v] {
		if !visited[next] {
			dfs(next, graph, visited)
		}
	}
}

func bfs(v int, graph [][]int, visited []bool) {
	queue := []int{v}
	visited[v] = true

	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		fmt.Print(tmp, " ")

		for _, next := range graph[tmp] {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
}
