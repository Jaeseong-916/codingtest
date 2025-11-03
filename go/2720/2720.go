package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var t, c int
	fmt.Fscanln(reader, &t)
	var res = make([]map[string]int, t)
	for i := range t {
		fmt.Fscanln(reader, &c)
		res[i] = make(map[string]int)
		res[i]["quater"] = c / 25
		c = c % 25
		res[i]["dime"] = c / 10
		c = c % 10
		res[i]["nickel"] = c / 5
		c = c % 5
		res[i]["penny"] = c
	}
	for i := range t {
		fmt.Printf("%d %d %d %d\n", res[i]["quater"], res[i]["dime"], res[i]["nickel"], res[i]["penny"])
	}
}
