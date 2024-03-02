package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	score := 0
	fp := flag.String("file", "problems.csv", "questions file path")
	shuffle := flag.Bool("shuffle", false, "shuffle questions")
	tl := flag.Int("limit", 30, "time limit in seconds")
	flag.Parse()
	qp := loadQuizFromFile(fp)
	if *shuffle {
		qp.shuffle()
	}
	t := time.NewTimer(time.Duration(*tl) * time.Second)

quizLoop:
	for _, p := range qp {
		fmt.Print(p.question, ": ")
		answerCh := make(chan string)
		go p.askQuestion(answerCh)
		select {
		case <-t.C:
			fmt.Println()
			break quizLoop
		case answer := <-answerCh:
			if answer == p.answer {
				score++
			}
		}
	}
	fmt.Printf("You have scored %d out of %d.\n", score, len(qp))
}
