package main

import (
	"encoding/csv"
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

func loadQuizFromFile(fp *string) Quiz {
	var qp Quiz
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
	return qp
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
