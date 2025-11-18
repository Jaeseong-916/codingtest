package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, s, m, l, xl, xxl, xxxl, t, p int
	var s_t, m_t, l_t, xl_t, xxl_t, xxxl_t int

	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &s, &m, &l, &xl, &xxl, &xxxl)
	fmt.Fscanln(reader, &t, &p)

	if s%t == 0 {
		s_t = s / t
	} else {
		s_t = (s / t) + 1
	}
	if m%t == 0 {
		m_t = m / t
	} else {
		m_t = (m / t) + 1
	}
	if l%t == 0 {
		l_t = l / t
	} else {
		l_t = (l / t) + 1
	}
	if xl%t == 0 {
		xl_t = xl / t
	} else {
		xl_t = (xl / t) + 1
	}
	if xxl%t == 0 {
		xxl_t = xxl / t
	} else {
		xxl_t = (xxl / t) + 1
	}
	if xxxl%t == 0 {
		xxxl_t = xxxl / t
	} else {
		xxxl_t = (xxxl / t) + 1
	}

	var res_shirt = s_t + m_t + l_t + xl_t + xxl_t + xxxl_t
	fmt.Println(res_shirt)
	fmt.Printf("%d %d\n", (n / p), (n % p))
}
