package problems

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

const questionTpl = "what %v, sir? "

type Problem struct {
	question string
	answer   string
}

func Handler(filename string, shuffle bool, answersChannel chan bool, endChannel chan bool) {
	n, problems, err := loadProblems(filename)
	if err != nil {
		fmt.Println(err)
		endChannel <- false
	}

	iterationIds := calculateIterationIds(n, shuffle)

	for _,i := range iterationIds {
		p := problems[i]

		printQuestion(p.question)

		userAnswer, err := getUserAnswer()
		if err != nil {
			answersChannel <- false
			continue
		}

		answersChannel <- userAnswer == p.answer
	}

	endChannel <- true
}

func loadProblems(filename string) (int, []*Problem, error) {
	var n int
	var problems []*Problem

	csvFile, err := os.Open(filename)
	if err != nil {
		return n, problems, err
	}

	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return n, problems, err
		}

		problems = append(problems, &Problem{line[0], strings.ToLower(strings.TrimSpace(line[1]))})
		n++
	}

	return n, problems, nil
}

func calculateIterationIds(n int, shuffle bool) (idsIteration []int){
	if shuffle {
		rand.Seed(time.Now().UnixNano())
		idsIteration = rand.Perm(n)
	} else {
		idsIteration = intRange(0, n, 1)
	}
	return
}

func intRange(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}

func printQuestion(question string) {
	fmt.Printf(questionTpl, question)
}

func getUserAnswer() (answer string, err error) {
	_, err = fmt.Scanf("%s", &answer)
	answer = strings.TrimSpace(answer)
	return
}