package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &m)

	// 인접 리스트로 그래프 구성
	graph := make([][]int, n+1)
	for range m {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	// BFS로 1번부터 탐색
	visited := make([]bool, n+1)
	queue := []int{1}
	visited[1] = true
	count := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, next := range graph[current] {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
				count++
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", count)
}
