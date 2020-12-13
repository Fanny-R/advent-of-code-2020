package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	bags := extractInput()
	fmt.Println(len(bags))

	searchedBags := make(map[string]bool)
	searchedBags["shiny gold"] = true
	var resultBags map[string]bool
	eligibleBags := make(map[string]bool)

	for len(searchedBags) > 0 {
		resultBags = make(map[string]bool)
		for bag, content := range bags {
			for searchedBag, _ := range searchedBags {
				if strings.Contains(content, searchedBag) {
					resultBags[bag] = true
					eligibleBags[bag] = true
				}
			}
		}
		searchedBags = resultBags
	}

	fmt.Println(len(eligibleBags))
}

func extractInput() map[string]string {
	fileBytes, err := ioutil.ReadFile("input/day7")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(string(fileBytes), "\n")
	bags := make(map[string]string)

	for _, line := range lines {
		splittedLine := strings.Split(line, "s contain ")
		bags[splittedLine[0]] = splittedLine[1]
	}

	return bags
}
