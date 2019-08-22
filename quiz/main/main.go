package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// problem is a struct to store question and correct answer
type problem struct {
	question string
	answer   string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// parseQuestions parses lines of CSV lines into set of problem
func parseQuestions(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func main() {
	// parsing input flags
	var csvFile string
	var timeLimit int
	flag.StringVar(&csvFile, "csv", "problems.csv", "a csv file in the format of (question, answer) (default problems.csv)")
	flag.IntVar(&timeLimit, "limit", 30, "time limit to answer each question (default 30 seconds)")
	flag.Parse()

	// reading the file name from args
	fileReader, err := os.Open(csvFile)
	defer fileReader.Close()
	check(err)

	// read the csv file
	csvReader := csv.NewReader(fileReader)
	lines, err := csvReader.ReadAll()
	check(err)

	// create problems out of the csv file
	problems := parseQuestions(lines)

	// create timer for each question
	timeUp := time.NewTimer(time.Duration(timeLimit) * time.Second)
	correct := 0

loop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		answerChannel := make(chan string)

		// ask question in a separate go-routine
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		// either time is up, or we get a correct answer in the channel
		select {
		case <-timeUp.C:
			fmt.Printf("\nYour time limit of %d seconds is up\n", timeLimit)
			break loop
		case answer := <-answerChannel:
			if answer == problem.answer {
				correct++
			}
		}
	}

	fmt.Printf("Your total score is: %d/%d \n", correct, len(problems))

}
