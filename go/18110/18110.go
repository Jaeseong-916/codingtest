package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	if n == 0 {
		fmt.Fprintf(writer, "%d\n", 0)
	} else {
		var avr_idx = int(math.Round(float64(n) * 0.15))
		var arr = []int{}

		for range n {
			var tmp int
			fmt.Fscanln(reader, &tmp)
			arr = append(arr, tmp)
		}

		// 오름차순 정리
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		// 15% 제거
		arr = arr[avr_idx:(len(arr) - avr_idx)]

		var sum int
		for _, val := range arr {
			sum += val
		}

		fmt.Fprintf(writer, "%d\n", int(math.Round(float64(sum)/float64(len(arr)))))
	}
}
