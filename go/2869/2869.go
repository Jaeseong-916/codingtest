package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b, v int
	fmt.Fscanln(reader, &a, &b, &v)
	if (v-a)%(a-b) != 0 {
		fmt.Println((v-a)/(a-b) + 2)
	} else {
		fmt.Println((v-a)/(a-b) + 1)
	}
}

// (v - a) % (a - b) != 0 결과 : (v - a) / (a - b) + 1
// (v - a) % (a - b) == 0 결과 : (v - a) / (a - b)

// func main() {
// 	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
// 	var a, b, v, current, day int
// 	fmt.Fscanln(reader, &a, &b, &v)
// 	for {
// 		day++
// 		current = current + a
// 		if current >= v {
// 			break
// 		}
// 		current = current - b
// 	}
// 	fmt.Println(day)
// }

// 0일떄는 v - a 나누기 a-b의 몫 +1 값
// (a-b) * ((v-a) / (a-b))
// 나머지 있으면, +2 나머지 없으면 +1
