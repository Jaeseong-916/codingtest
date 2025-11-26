package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Ch(input []string) string {
	var res string
	var idx, tmp int
	for i := range input {
		if num, err := strconv.Atoi(input[i]); err == nil {
			tmp = num
			idx = i
		}
	}
	switch idx {
	case 0:
		if (tmp+3)%3 == 0 && (tmp+3)%5 == 0 {
			res = "FizzBuzz"
		} else if (tmp+3)%3 == 0 && (tmp+3)%5 != 0 {
			res = "Fizz"
		} else if (tmp+3)%3 != 0 && (tmp+3)%5 == 0 {
			res = "Buzz"
		} else {
			res = strconv.Itoa(tmp + 3)
		}
	case 1:
		if (tmp+2)%3 == 0 && (tmp+2)%5 == 0 {
			res = "FizzBuzz"
		} else if (tmp+2)%3 == 0 && (tmp+2)%5 != 0 {
			res = "Fizz"
		} else if (tmp+2)%3 != 0 && (tmp+2)%5 == 0 {
			res = "Buzz"
		} else {
			res = strconv.Itoa(tmp + 2)
		}
	case 2:
		if (tmp+1)%3 == 0 && (tmp+1)%5 == 0 {
			res = "FizzBuzz"
		} else if (tmp+1)%3 == 0 && (tmp+1)%5 != 0 {
			res = "Fizz"
		} else if (tmp+1)%3 != 0 && (tmp+1)%5 == 0 {
			res = "Buzz"
		} else {
			res = strconv.Itoa(tmp + 1)
		}
	}
	return res
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input = make([]string, 3)
	for i := range input {
		fmt.Fscanln(reader, &input[i])
	}

	fmt.Fprintf(writer, "%s\n", Ch(input))

}
