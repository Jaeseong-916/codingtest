package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, i int
	var bee = make([]int, 1)
	var tmp int = 1

	for {
		if tmp > 1000000000 || tmp < 1 {
			break
		}
		tmp = tmp + 6*(i+1)
		bee = append(bee, tmp)
		i++
	}

	fmt.Fscanln(reader, &n)

	for j := range len(bee) - 1 {
		if n == 1 {
			fmt.Println("1")
			break
		} else if n > bee[j] && n <= bee[j+1] {
			fmt.Println(j + 2)
			break
		}
	}

}
