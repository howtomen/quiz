package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	score     int
	file_path string
	res       string
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

	for _, prob := range problems {
		fmt.Println(prob.q)
		fmt.Scanln(&res)
		if normalize(res) == normalize(prob.a) {
			score++
		}
	}
}
func main() {
	flag.StringVar(&file_path, "path", "problems.csv", "Location of CSV with Questions and answers in question,answer format") //String(name,defaultValue,usage)
	flag.Parse()

	file, error := os.Open(file_path)
	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()

	problems := parseFile(file)

	takeQuiz(problems)

	fmt.Printf("You got %d correct out of %d", score, len(problems))
}
