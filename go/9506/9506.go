package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var input int
	var input_arr []int
	for {
		fmt.Fscanln(reader, &input)
		if input == -1 {
			break
		}
		input_arr = append(input_arr, input)
	}
	for i := range input_arr {
		var sum int = 0
		var res []int
		for j := 1; j < input_arr[i]; j++ {
			if input_arr[i]%j == 0 {
				sum = sum + j
				res = append(res, j)
			}
		}
		if input_arr[i] == sum {
			fmt.Printf("%d = ", input_arr[i])
			for j := range res {
				fmt.Printf("%d", res[j])
				if j != len(res)-1 {
					fmt.Print(" + ")
				}
			}
			fmt.Println("")
		} else {
			fmt.Printf("%d is NOT perfect.\n", input_arr[i])
		}
	}
}
