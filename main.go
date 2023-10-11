package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	score     int
	file_path string
	res       string
	timeLimit int
)

type problem struct {
	q string
	a string
}

func normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func parseFile(file *os.File) []problem {
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i].q = line[0]
		problems[i].a = line[1]
	}

	return problems
}

func takeQuiz(problems []problem) {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for i, prob := range problems {
		fmt.Printf("Problem %d: %s = ", i+1, prob.q)
		ansCh := make(chan string)
		go func() {
			fmt.Scanln(&res)
			ansCh <- res
		}()
		select {
		case <-timer.C:
			return
		case answer := <-ansCh:
			if normalize(answer) == normalize(prob.a) {
				score++
			}
		}
	}

}
func main() {
	flag.StringVar(&file_path, "path", "problems.csv", "Location of CSV with Questions and answers in question,answer format") //String(name,defaultValue,usage)
	flag.IntVar(&timeLimit, "timer", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, error := os.Open(file_path)
	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()

	problems := parseFile(file)

	takeQuiz(problems)

	fmt.Printf("\nYou got %d correct out of %d", score, len(problems))
}
