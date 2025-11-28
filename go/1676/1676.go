package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func factorial(v *big.Int, n int64) *big.Int {
	var t *big.Int = big.NewInt(1)
	for i := range n {
		t = big.NewInt(i + 1)
		v = v.Mul(v, t)
	}
	return v
}

func cntZero(v *big.Int) int {
	strn := v.String()
	var cnt int
	for i := len(strn) - 1; i >= 0; i-- {
		if strn[i] != byte('0') {
			break
		}
		cnt++
	}
	return cnt
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int64
	var v *big.Int

	fmt.Fscanln(reader, &n)

	v = big.NewInt(1)
	fmt.Fprintf(writer, "%d\n", cntZero(factorial(v, n)))
}
