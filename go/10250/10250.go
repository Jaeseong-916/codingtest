package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var t, h, w, n int

	fmt.Fscanln(reader, &t)
	var res []string

	for range t {
		fmt.Fscanln(reader, &h, &w, &n)
		res = append(res, room_num(h, w, n))
	}
}

func room_num(h int, w int, n int) string {
	if n == 1 {
		return "101"
	}
	var num []int
	var cnt int = 1
	for i := 0; i < h; i++ { // 처음부터 h까지 1씩 더하면서 진행.
		cnt = cnt + i
		num = append(num, cnt)
	}
	for i := h; i < h-w; i++ { // h 부터 w-h까지 h씩 더하면서 진행.
		cnt = cnt + h
		num = append(num, cnt)
	}
	for i := h - 1; i < 0; i-- { // w 이후 h만큼 1씩 감소하면서 진행.
		cnt = cnt + i
		num = append(num, cnt)
	}
	for i := range num {
		if num[i] <= n {
			if num[i]%n == 0 {
				if i < 10 {
					return "10" + strconv.Itoa(i)
				} else {
					return "1" + strconv.Itoa(i)
				}
			} else {

			}
		}
	}
}
