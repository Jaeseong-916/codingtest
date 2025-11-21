package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var t, k, n int

	fmt.Fscanln(reader, &t)
	var res = make([]int, t)

	for i := range res {
		fmt.Fscanln(reader, &k)
		fmt.Fscanln(reader, &n)
		var mp = make([][]int, k+1)
		if n == 1 {
			res[i] = 1
			continue
		}
		if k == 0 {
			for j := range n {
				res[i] = res[i] + j + 1
			}
			continue
		}
		for j := range k + 1 {
			mp[j] = make([]int, n)
			for l := range n {
				if j == 0 {
					mp[j][l] = mp[j][l] + l + 1
				} else if l == 0 {
					mp[j][l] = 1
				} else {
					mp[j][l] = mp[j-1][l] + mp[j][l-1]
				}
			}
		}
		res[i] = mp[k][n-1]
	}
	for i := range res {
		fmt.Println(res[i])
	}
}
