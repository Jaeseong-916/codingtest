package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var l int
	var res int64
	var input string
	var pow int64 = 1

	fmt.Fscanln(reader, &l)
	fmt.Fscanln(reader, &input)

	for i := range l {
		res = (res + (int64(input[i]-96) * pow % 1234567891)) % 1234567891
		pow = (pow * 31) % 1234567891
	}
	fmt.Println(res)
}
