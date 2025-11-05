package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var x int
	var arr = []int{1}
	var val, i = 1, 1

	for {
		if val > 10000000 {
			break
		}
		val = val + i + 1
		arr = append(arr, val)
		i++
	}

	fmt.Fscanln(reader, &x)

	for j := range arr {
		if x == 1 {
			fmt.Println("1/1")
			break
		}
		if j > 0 {
			if j%2 == 1 {
				if arr[j-1] < x && arr[j] >= x {
					fmt.Printf("%d/%d\n", x-arr[j-1], (j+2)-(x-arr[j-1])) // 위에서 아래로 (홀수)
					break
				}
			} else {
				if arr[j-1] < x && arr[j] >= x {
					fmt.Printf("%d/%d\n", (j+2)-(x-arr[j-1]), x-arr[j-1]) // 아래에서 위로 (짝수)
					break
				}
			}
		}
	}
}
