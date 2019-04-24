package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
	"encoding/csv"
)

// QAndA - struct to transform the questions and answers to JSON
type QAndA struct {
		Question string
		Answer string 
	}

func main () {
	csvFileName := flag.String("csv", "problems.csv", "a csv file containing questions and solutions in the form  'question,answer'")
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
	correct := 0
	wrong := 0
	for index, question := range qAndAs {
		var answer string
		fmt.Printf("Question number %d: %s = ", index+1, question.Question)
		_, error := fmt.Scan(&answer)
		if error != nil {
			fmt.Printf("Couldn't read answer")
		}
		if answer == question.Answer {
			correct++
		} else {
			wrong++
		}
	}
	fmt.Printf("You scored %d/%d\n", correct, len(qAndAs))
}

func parseLines(lines [][]string) []QAndA {
	ret := make([]QAndA, len(lines))
	for i, line := range lines {
		ret[i] = QAndA{
			Question: line[0],
			Answer: strings.TrimSpace(line[1]),
		}
	}
	return ret
}