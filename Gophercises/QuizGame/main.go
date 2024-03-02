package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Quiz []Problem

type Problem struct {
	question string
	answer   string
}

func main() {
	var qp Quiz
	score := 0
	fp := flag.String("file", "problems.csv", "questions file path")
	shuffle := flag.Bool("shuffle", false, "shuffle questions")
	tl := flag.Int("limit", 30, "time limit in seconds")
	flag.Parse()
	f, err := os.Open(*fp)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	for _, p := range csvData {
		qp = append(qp, Problem{question: p[0], answer: strings.TrimSpace(p[1])})
	}
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

func (p Problem) askQuestion(c chan string) {
	var ans string
	fmt.Scanf("%s\n", &ans)
	c <- ans
}

func (q Quiz) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range q {
		newPos := r.Intn(len(q) - 1)
		q[i], q[newPos] = q[newPos], q[i]
	}
}
