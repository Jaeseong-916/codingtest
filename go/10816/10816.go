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
	var cardMap = make(map[int]int)

	// N 입력
	fmt.Fscanln(reader, &n)

	// N개의 숫자 카드 입력
	cardsLine, _ := reader.ReadString('\n')
	cardsLine = strings.TrimSpace(cardsLine)
	cardStrs := strings.Split(cardsLine, " ")

	for _, cardStr := range cardStrs {
		num, _ := strconv.Atoi(cardStr)
		cardMap[num]++
	}

	// M 입력
	fmt.Fscanln(reader, &m)

	// M개의 쿼리 입력
	queryLine, _ := reader.ReadString('\n')
	queryLine = strings.TrimSpace(queryLine)
	queryStrs := strings.Split(queryLine, " ")

	// 결과 출력
	for i, queryStr := range queryStrs {
		num, _ := strconv.Atoi(queryStr)
		if i > 0 {
			fmt.Fprintf(writer, " ")
		}
		fmt.Fprintf(writer, "%d", cardMap[num])
	}
	fmt.Fprintf(writer, "\n")
}
