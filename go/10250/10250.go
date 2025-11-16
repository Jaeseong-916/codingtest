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
	var res []int

	for range t {
		fmt.Fscanln(reader, &h, &w, &n)
		res = append(res, room_num(h, n))
	}
	for i := range t {
		fmt.Println(res[i])
	}
}

func room_num(h int, n int) int {
	var room, floor, res int
	room = (n / h) + 1
	floor = n % h
	if n%h == 0 {
		room = (n / h)
		floor = h
	}
	if room > 9 {
		res, _ = strconv.Atoi(strconv.Itoa(floor) + strconv.Itoa(room))
	} else {
		res, _ = strconv.Atoi(strconv.Itoa(floor) + "0" + strconv.Itoa(room))
	}
	return res
}

// 문제 해석을 잘못해서 나온 결과
// func room_num(h int, w int, n int) string {
// 	if n == 1 {
// 		return "101"
// 	}
// 	var num []int
// 	var cnt int = 1
// 	for i := 0; i < h; i++ { // 처음부터 h까지 1씩 더하면서 진행.
// 		cnt = cnt + i
// 		num = append(num, cnt)
// 	}
// 	for i := h; i < w; i++ { // h 부터 w-h까지 h씩 더하면서 진행.
// 		cnt = cnt + h
// 		num = append(num, cnt)
// 	}
// 	for i := h - 1; i > 0; i-- { // w 이후 h만큼 1씩 감소하면서 진행.
// 		cnt = cnt + i
// 		num = append(num, cnt)
// 	}
// 	// i == 4, i - 1 == 3, num[3] == 7
// 	// 4 - 7
// 	for i := range num {
// 		if num[i] >= n {
// 			if num[i]%n == 0 {
// 				if i < 10 {
// 					fmt.Println("1")
// 					return "10" + strconv.Itoa(i)
// 				} else {
// 					fmt.Println("2")
// 					return "1" + strconv.Itoa(i)
// 				}
// 			} else if num[i]%n < 10 {
// 				if i < 9 {
// 					fmt.Println("3")
// 					return strconv.Itoa(n%num[i-1]) + "0" + strconv.Itoa(i-(n%num[i-1]))
// 				} else {
// 					fmt.Println("4")
// 					return strconv.Itoa(num[i-1]%n) + strconv.Itoa(i-num[i-1]%n)
// 				}
// 			} else {
// 				if i < 9 {
// 					fmt.Println("5")
// 					return strconv.Itoa(num[i-1]%n) + "0" + strconv.Itoa(i-num[i-1]%n)
// 				} else {
// 					fmt.Println("6")
// 					return strconv.Itoa(num[i-1]%n) + strconv.Itoa(i-num[i-1]%n)
// 				}
// 			}
// 		}
// 	}
// 	return ""
// }
