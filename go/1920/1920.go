package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int

	fmt.Fscanln(reader, &n)
	var n_map = make(map[int]int)

	var line string
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	// 공백으로 분리
	numbers := strings.FieldsSeq(line)

	// 각 문자열을 정수로 변환하여 map에 저장
	for numStr := range numbers {
		num, _ := strconv.Atoi(numStr)
		n_map[num]++
	}

	var m int
	var m_arr []int
	fmt.Fscanln(reader, &m)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	nums := strings.FieldsSeq(line)

	for numStr := range nums {
		num, _ := strconv.Atoi(numStr)
		m_arr = append(m_arr, num)
	}

	for i := range m_arr {
		if n_map[m_arr[i]] != 0 {
			fmt.Fprintf(writer, "%d\n", 1)
		} else {
			fmt.Fprintf(writer, "%d\n", 0)
		}
	}
}
