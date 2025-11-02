package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	input, _ = reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	input = strings.Replace(input, "c=", " ", -1)
	input = strings.Replace(input, "c-", " ", -1)
	input = strings.Replace(input, "dz=", " ", -1)
	input = strings.Replace(input, "d-", " ", -1)
	input = strings.Replace(input, "lj", " ", -1)
	input = strings.Replace(input, "nj", " ", -1)
	input = strings.Replace(input, "s=", " ", -1)
	input = strings.Replace(input, "z=", " ", -1)

	fmt.Println(len(input))
}

// func main() {
// 	var input string
// 	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// 	input, _ = reader.ReadString('\n')
// 	input = strings.TrimSuffix(input, "\n")

// 	var res int = 0
// 	for idx, i := range input {
// 		res++
// 		if idx != 0 {
// 			switch idx != 0 {
// 			case string(i) == "-":
// 				res--
// 			case string(i) == "=":
// 				if idx > 1 {
// 					if string(input[idx-2]) == "d" {
// 						res = res - 2
// 					} else {
// 						res--
// 					}
// 				} else {
// 					res--
// 				}
// 			case string(i) == "j":
// 				if string(input[idx-1]) == "l" || string(input[idx-1]) == "n" {
// 					res--
// 				}
// 			}
// 		}
// 		fmt.Println(res)
// 	}
// 	fmt.Print(res)
// }
