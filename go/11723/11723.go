package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func add(arr []int, n int) []int {
	var exist = false
	if slices.Contains(arr, n) {
		exist = true
	}
	if !exist {
		arr = append(arr, n)
	}
	return arr
}

func remove(arr []int, n int) []int {
	for i, num := range arr {
		if num == n {
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	return arr
}

func check(arr []int, n int) bool {
	var res = false
	for _, num := range arr {
		if num == n {
			res = true
			break
		}
	}
	return res
}

func toggle(arr []int, n int) []int {
	var exist = false
	for _, num := range arr {
		if num == n {
			exist = true
			break
		}
	}
	if exist {
		arr = remove(arr, n)
	} else {
		arr = add(arr, n)
	}
	return arr
}

func all(arr []int) []int {
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	return arr
}

func empty(arr []int) []int {
	arr = []int{}
	return arr
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var m int
	fmt.Fscanln(reader, &m)

	var s = []int{}

	for range m {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		// 공백으로 분리
		parts := strings.Fields(line)

		if len(parts) > 0 {
			// command에 명령어 입력
			command := parts[0]

			if command == "add" && len(parts) == 2 {
				// add
				val, _ := strconv.Atoi(parts[1])
				s = add(s, val)
			} else if command == "remove" && len(parts) == 2 {
				// remove
				val, _ := strconv.Atoi(parts[1])
				s = remove(s, val)
			} else if command == "check" && len(parts) == 2 {
				// remove
				val, _ := strconv.Atoi(parts[1])
				if check(s, val) {
					fmt.Fprintf(writer, "1\n")
				} else {
					fmt.Fprintf(writer, "0\n")
				}
			} else if command == "toggle" && len(parts) == 2 {
				// remove
				val, _ := strconv.Atoi(parts[1])
				s = toggle(s, val)
			} else if command == "all" {
				// all
				s = all(s)
			} else if command == "empty" {
				// empty
				s = empty(s)
			}
		}
	}
}
