package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Srt(input []string) []string {
	// 중복 제거를 위해 map 사용
	unique := make(map[string]bool)
	for _, word := range input {
		unique[word] = true
	}

	// map에서 slice로 변환
	var res []string
	for word := range unique {
		res = append(res, word)
	}

	// 정렬: 1. 길이순, 2. 사전순
	sort.Slice(res, func(i, j int) bool {
		if len(res[i]) == len(res[j]) {
			return res[i] < res[j] // 길이가 같으면 사전순
		}
		return len(res[i]) < len(res[j]) // 길이가 짧은 순
	})

	return res
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var input = make([]string, n)

	for i := range input {
		fmt.Fscanln(reader, &input[i])
	}

	res := Srt(input)

	for i := range res {
		fmt.Fprintln(writer, res[i])
	}
}
