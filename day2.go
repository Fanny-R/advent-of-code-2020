package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	// input := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
	input := extractInput()

	valid := 0
	for _, line := range input {
		lineData := strings.Split(line, " ")

		nt := strings.Split(lineData[0], "-")
		lowest, _ := strconv.Atoi(nt[0])
		highest, _ := strconv.Atoi(nt[1])

		letter := strings.TrimSuffix(lineData[1], ":")
		password := lineData[2]

		count := strings.Count(password, letter)

		if count >= lowest && count <= highest {
			valid++
		}
	}
	fmt.Println(valid)
}

func extractInput() []string {
	fileBytes, err := ioutil.ReadFile("input/day2")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(fileBytes), "\n")
}
