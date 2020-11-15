package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

// question struct stores a single question and its corresponding answer.
type question struct {
	q, a string
}

type score int

// check handles a potential error.
// It stops execution of the program ("panics") if an error has happened.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// questions reads in questions and corresponding answers from a CSV file into a slice of question structs.
func questions() []question {
	f, err := os.Open("quiz-questions.csv")
	check(err)
	reader := csv.NewReader(f)
	table, err := reader.ReadAll()
	check(err)
	var questions []question
	for _, row := range table {
		questions = append(questions, question{q: row[0], a: row[1]})
	}
	return questions
}

// ask asks a question and returns an updated score depending on the answer.
//ask is goroutine; blocks and waits for input while main continues to work
func ask(s score, question question, newScore chan score) {
	fmt.Println(question.q)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter answer: ")
	scanner.Scan()
	text := scanner.Text()
	if strings.Compare(text, question.a) == 0 {
		fmt.Println("Correct!")
		s++
	} else {
		fmt.Println("Incorrect :-(")
	}
	newScore <- s
}

// func main() {
// 	s := score(0)
// 	qs := questions()
// 	newScore := make(chan score)
// 	for _, q := range qs {
// 		go ask(s, q, newScore)
// 		s = <-newScore
// 	}
// 	fmt.Println("Final score", s)
// }

//make quiz run for 5s only
func main() {
	s := score(0)
	qs := questions()
	timeOver := time.After(5 * time.Second)
	scoreChan := make(chan score)
	for _, q := range qs {
		go ask(s, q, scoreChan)
		select {
		case newScore := <-scoreChan:
			s = newScore
		case <-timeOver:
			fmt.Println("Timed out, final score", s)
			return
		}
	}
	fmt.Println("Final score", s)
}
