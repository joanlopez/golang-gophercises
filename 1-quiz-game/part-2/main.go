package main

import (
	"flag"
	"fmt"
	"github.com/joanlopez/golang-gophercises/1-quiz-game/part-2/problems"
	"github.com/joanlopez/golang-gophercises/1-quiz-game/part-2/timeout"
	"os"
)

const defaultTimeout = 30
const defaultShuffle = false
const problemsFile = "1-quiz-game/part-2/problems.csv"

const beginTxt = "Press enter to start the quiz..."
const endTxt = "\n\n%v correct answers of %v\n"

func main() {
	fileNamePtr := flag.String("file", problemsFile, "problems file name")
	timeoutPtr := flag.Int("timeout", defaultTimeout, "challenge's timeout (seconds)")
	shufflePtr := flag.Bool("shuffle", defaultShuffle, "shuffle problems collection")
	flag.Parse()

	totalAnswers := 0
	correctAnswers := 0

	answersChannel := make(chan bool)
	endChannel := make(chan bool)

	fmt.Println(beginTxt)
	_, _ = fmt.Scanln()

	go timeout.Handler(*timeoutPtr, endChannel)
	go problems.Handler(*fileNamePtr, *shufflePtr, answersChannel, endChannel)

	for {
		select {
		case answer := <-answersChannel:
			totalAnswers++
			if answer {
				correctAnswers++
			}
		case ok := <-endChannel:
			close(answersChannel)
			close(endChannel)
			terminate(ok, correctAnswers, totalAnswers)
		}
	}
}

func terminate(ok bool, correctAnswers, totalAnswers int) {
	fmt.Printf(endTxt, correctAnswers, totalAnswers)
	if ok {
		os.Exit(0)
	}
	os.Exit(1)

}
