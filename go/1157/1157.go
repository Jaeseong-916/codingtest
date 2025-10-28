package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input, res string
	var maxVal int

	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.ToUpper(input)

	cnt := make(map[string]int)
	for _, c := range input {
		if string(c) != " " {
			cnt[string(c)]++
		}
	}

	for char, occur := range cnt {
		if occur > maxVal {
			maxVal = occur
			res = string(char)
		} else if occur == maxVal {
			res = "?"
		}
	}
	fmt.Print(res)
}

// 틀린 풀이

// type mp struct {
// 	key   rune
// 	value int
// }

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	input, _ := reader.ReadString('\n')

// 	m := make(map[rune]int)

// 	for _, c := range input {
// 		if c > 91 {
// 			c = c - 32
// 		}
// 		m[c] = m[c] + 1
// 	}

// 	delete(m, rune(10))

// 	var mps []mp

// 	for k, v := range m {
// 		mps = append(mps, mp{k, v})
// 	}

// 	if len(m) > 1 {
// 		sort.Slice(mps, func(i, j int) bool {
// 			return mps[i].value < mps[j].value
// 		})

// 		if mps[len(mps)-1] == mps[len(mps)-2] {
// 			fmt.Print("?")
// 		} else {
// 			fmt.Print(string(mps[len(mps)-1].key))
// 		}
// 	} else {
// 		fmt.Print(string(mps[0].key))
// 	}
// }
