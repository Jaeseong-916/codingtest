package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var k, n int
	fmt.Fscanln(reader, &k, &n)

	var length_arr []int
	for range k {
		var tmp int
		fmt.Fscanln(reader, &tmp)
		length_arr = append(length_arr, tmp)
	}
	sort.Ints(length_arr)

	// 이진 탐색
	left := 1
	right := length_arr[k-1]
	var result int

	for left <= right {
		mid := (left + right) / 2
		var count int
		for j := 0; j < k; j++ {
			count += length_arr[j] / mid
		}
		if count >= n {
			// 더 큰 길이 시도
			result = mid
			left = mid + 1
		} else {
			// 더 작은 길이로 시도
			right = mid - 1
		}
	}

	fmt.Fprintf(writer, "%d\n", result)
}
