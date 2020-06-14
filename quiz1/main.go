package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// timeout := make(chan time.Time, 1)
	file, err := os.Open("quizlist.csv")
	if err != nil {
		fmt.Println("error opening the file. ", err)
		os.Exit(1)
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	quizList := makeQuizList(lines)

	// todo make the timeLimit as parameter
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("your timer has started.")

	var answerChan = make(chan string)

	correctAnswers := 0

loop:
	for i, quiz := range quizList {

		fmt.Printf("problem #%d:  %s= ", i+1, quiz.question)
		readAnswer(answerChan)
		select {
		case <-timer.C:
			fmt.Printf("\ntimes up...\n")
			break loop
		case answer := <-answerChan:
			if answer == quiz.answer {
				correctAnswers++
			}
		}
	}
	fmt.Printf("Your score is %d out of %d. \n", correctAnswers, len(quizList))
}

func makeQuizList(lines [][]string) []quiz {
	quizList := make([]quiz, len(lines))
	for i, line := range lines {
		quizList[i] = quiz{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return quizList
}
func readAnswer(answerChan chan<- string) {
	go func() {
		var answer string
		fmt.Scanf("%s", &answer)
		answerChan <- answer
	}()
}

type quiz struct {
	question string
	answer   string
}
