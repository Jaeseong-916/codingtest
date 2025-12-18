package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func avg(arr []int) int {
	var res, sum int
	for _, i := range arr {
		sum += i
	}
	res = int(math.Round(float64(sum) / float64(len(arr))))
	return res
}

func mid(arr []int) int {
	return arr[(len(arr) / 2)]
}

func frequency(arr []int) int {
	var mp = make(map[int]int)
	for _, i := range arr {
		mp[i]++
	}
	var max int
	for i := -4000; i <= 4000; i++ {
		if mp[i] > max {
			max = mp[i]
		}
	}
	var res = []int{}
	for i := -4000; i <= 4000; i++ {
		if mp[i] == max {
			res = append(res, i)
		}
	}

	if len(res) < 2 {
		return res[0]
	} else {
		return res[1]
	}
}

func ran(arr []int) int {
	return arr[len(arr)-1] - arr[0]
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var arr = []int{}
	for range n {
		var tmp int
		fmt.Fscanln(reader, &tmp)
		arr = append(arr, tmp)
	}

	slices.Sort(arr)

	fmt.Fprintf(writer, "%d\n%d\n%d\n%d\n", avg(arr), mid(arr), frequency(arr), ran(arr))
}
