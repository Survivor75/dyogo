package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"log"
	"strconv"
)

type stack []int

func (s *stack) push(val int) {
	*s = append(*s, val)
}

func (s *stack) pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

func evaluator(question string, answer int) bool {
	fmt.Println(question, answer)
	var stk stack = make([]int, 0)
	var operator string
	for _, element := range question {
		if val, err := strconv.Atoi(string(element)); err == nil{
			stk.push(val)
		} else {
			operator = string(element)
		}
	}
	operand1 := stk.peek()
	stk.pop()
	operand2 := stk.peek()
	stk.pop()
	switch operator {
		case "+":
			return operand1 + operand2 == answer
		case "-":
			return operand1 - operand2 == answer
		case "*":
			return operand1 * operand2 == answer
		case "/":
			return operand1 / operand2 == answer
		default:
			return false
	}
}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var record []string
	for record, err = reader.Read(); err != io.EOF; record, err = reader.Read(){
		if err != nil {
			log.Fatal(err)
		}
		answer, _ := strconv.Atoi(string(record[1]))
		fmt.Println(evaluator(record[0], answer))
	}
}
