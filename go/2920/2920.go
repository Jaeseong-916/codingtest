package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var cnt_asc, cnt_des int
	var input = make([]int, 8)

	for i := range 8 {
		fmt.Fscan(reader, &input[i])
	}

	for i := range 7 {
		if input[0] == 1 && input[7] == 8 {
			if input[i] == input[i+1]-1 {
				cnt_asc++
			} else {
				break
			}
		}
		if input[0] == 8 && input[7] == 1 {
			if input[i] == input[i+1]+1 {
				cnt_des++
			} else {
				break
			}
		}
	}

	if cnt_asc == 7 {
		fmt.Println("ascending")
	} else if cnt_des == 7 {
		fmt.Println("descending")
	} else {
		fmt.Println("mixed")
	}
}
