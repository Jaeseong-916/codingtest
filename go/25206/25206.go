package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var lec, grade string
	var score float32

	var total_score, total_score_grade float32

	for i := 0; i < 20; i++ {
		fmt.Fscanln(reader, &lec, &score, &grade)

		switch grade != " " {
		case grade == "A+":
			total_score_grade = total_score_grade + score*4.5
			total_score = total_score + score
		case grade == "A0":
			total_score_grade = total_score_grade + score*4.0
			total_score = total_score + score
		case grade == "B+":
			total_score_grade = total_score_grade + score*3.5
			total_score = total_score + score
		case grade == "B0":
			total_score_grade = total_score_grade + score*3.0
			total_score = total_score + score
		case grade == "C+":
			total_score_grade = total_score_grade + score*2.5
			total_score = total_score + score
		case grade == "C0":
			total_score_grade = total_score_grade + score*2.0
			total_score = total_score + score
		case grade == "D+":
			total_score_grade = total_score_grade + score*1.5
			total_score = total_score + score
		case grade == "D0":
			total_score_grade = total_score_grade + score*1.0
			total_score = total_score + score
		case grade == "F":
			total_score_grade = total_score_grade + score*0.0
			total_score = total_score + score
		case grade == "P":
			continue
		}
	}
	fmt.Printf("%f\n", total_score_grade/total_score)
}
