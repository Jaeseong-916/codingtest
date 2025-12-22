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

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var dict_str = make(map[string]int)
	var dict_num = make(map[int]string)
	for i := 1; i <= n; i++ {
		var tmp, _ = reader.ReadString('\n')
		tmp = strings.TrimSpace(tmp)
		dict_str[tmp] = i
		dict_num[i] = tmp
	}

	var res = []string{}
	for range m {
		var tmp, _ = reader.ReadString('\n')
		tmp = strings.TrimSpace(tmp)
		num, err := strconv.Atoi(tmp)
		if err != nil {
			// 문자 입력받았을 경우
			res = append(res, strconv.Itoa(dict_str[tmp]))
		} else {
			// 숫자 입력받았을 경우
			res = append(res, dict_num[num])
		}
	}

	for _, str := range res {
		fmt.Fprintf(writer, "%s\n", str)
	}
}

// array 시간초과

// var dict = []string{}
// for range n {
// 	var tmp, _ = reader.ReadString('\n')
// 	dict = append(dict, strings.TrimSpace(tmp))
// }

// var res = []string{}
// for range m {
// 	var tmp, _ = reader.ReadString('\n')
// 	tmp = strings.TrimSpace(tmp)
// 	num, err := strconv.Atoi(tmp)
// 	if err != nil {
// 		res = append(res, strconv.Itoa(slices.Index(dict, tmp)+1))
// 	} else {
// 		res = append(res, dict[num-1])
// 	}
// }

// for _, str := range res {
// 	fmt.Fprintf(writer, "%s\n", str)
// }
