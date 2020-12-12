package main

import (
	"fmt"
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

		answersInOneGroup := make(map[string]int)
		for _, personAnswers := range personsAnswers {
			for _, personAnswer := range personAnswers {
				answersInOneGroup[string(personAnswer)] = 0
			}
		}
		result += len(answersInOneGroup)
	}

	fmt.Println("result")
	fmt.Println(result)
}
