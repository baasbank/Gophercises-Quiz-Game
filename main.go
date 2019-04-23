package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"encoding/json"
	"encoding/csv"
	"io"
	"log"
)

// QAndA - struct to transform the questions and answers to JSON
type QAndA struct {
		Question string `json:"question"`
		Answer string `json:"answer"`
	}

func main () {
	csvFileName := flag.String("csv", "problems.csv", "a csv file containing questions and solutions in the form  'question,answer'")
	flag.Parse()
	
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println("Failed to open file")
	}
	reader := csv.NewReader(bufio.NewReader(file))
	var qAndAnswers []QAndA
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		qAndAnswers = append(qAndAnswers, QAndA{
			Question: line[0],
			Answer: line[1],
		})
		}
		questionsAndAnswers, _ := json.Marshal(qAndAnswers)
		fmt.Println(string(questionsAndAnswers))
}
