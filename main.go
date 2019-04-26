package main

import (
	"encoding/csv"
	"fmt"
	"flag"
	"os"
	"strings"
	"time"
	
)

// QAndA - struct to transform the questions and answers to JSON
type QAndA struct {
		question string
		answer string 
	}

func main () {
	csvFileName := flag.String("csv", "problems.csv", "a csv file containing questions and solutions in the form  'question,answer'")
	timeLimit := flag.Int("time", 30, "a number denoting time limit for each question in seconds")
	flag.Parse()

	
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", *csvFileName)
	}
	reader := csv.NewReader(file)
	lines, error := reader.ReadAll()
	if error != nil {
		fmt.Printf("Failed to parse %s\n", *csvFileName)
	}
	qAndAs := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	questionsLabel:
	for index, csvRow := range qAndAs {
		fmt.Printf("Question number %d: %s = ", index+1, csvRow.question)
		inputChannel := make(chan string)
		go func() {
			var input string
			_, error := fmt.Scan(&input)
			if error != nil {
				fmt.Printf("Couldn't read answer")
			}
			inputChannel <- input
		}()
		select {
		case <-timer.C:
			fmt.Println("\nYou exceeded your time limit")
			break questionsLabel
		case input := <-inputChannel:
			if strings.ToLower(input) == strings.ToLower(csvRow.answer) {
			correct++
			fmt.Println("\u2713")
			} else {
				fmt.Println("\u2717")
			}
		}
	}
	fmt.Printf("You scored %d/%d\n", correct, len(qAndAs))
}

func parseLines(lines [][]string) []QAndA {
	ret := make([]QAndA, len(lines))
	for i, line := range lines {
		ret[i] = QAndA{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return ret
}