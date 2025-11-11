package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b, c int
	var res []string
	for {
		fmt.Fscanln(reader, &a, &b, &c)
		if a == 0 && b == 0 && c == 0 {
			break
		}
		if a == b && b == c {
			res = append(res, "Equilateral")
			continue
		}
		if a >= b && a >= c {
			if b+c <= a {
				res = append(res, "Invalid")
				continue
			} else if a == b || a == c || b == c {
				res = append(res, "Isosceles")
				continue
			} else {
				res = append(res, "Scalene")
				continue
			}
		} else if b >= a && b >= c {
			if a+c <= b {
				res = append(res, "Invalid")
				continue
			} else if b == c || b == a || a == c {
				res = append(res, "Isosceles")
				continue
			} else {
				res = append(res, "Scalene")
				continue
			}
		} else if c >= a && c >= b {
			if a+b <= c {
				res = append(res, "Invalid")
				continue
			} else if c == a || c == b || a == b {
				res = append(res, "Isosceles")
				continue
			} else {
				res = append(res, "Scalene")
				continue
			}
		}
	}
	for i := range res {
		fmt.Println(res[i])
	}
}
