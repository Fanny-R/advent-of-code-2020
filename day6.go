package main

import (
	"fmt"
	"github.com/juliangruber/go-intersect"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileBytes, err := ioutil.ReadFile("input/day6")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	groups := strings.Split(string(fileBytes), "\n\n")
	result := 0
	for _, group := range groups {
		personsAnswers := strings.Split(group, "\n")
		var commonAnswers []interface{}
		for key, personAnswers := range personsAnswers {
			answerInterface := make([]interface{}, len(personAnswers))
			for i, answer := range personAnswers {
				answerInterface[i] = answer
			}
			if key == 0 {
				commonAnswers = answerInterface
				continue
			} else {
				commonAnswers = intersect.Simple(commonAnswers, answerInterface)
			}
		}
		result += len(commonAnswers)
	}

	fmt.Println("result")
	fmt.Println(result)
}
