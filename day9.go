package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := extractInput()
	preamble := 25

	for i := preamble; i < len(input); i++ {
		valid := false
		for j := i - preamble; j < i-1; j++ {
			for k := i - preamble + 1; k < i; k++ {
				if input[i] == input[j]+input[k] {
					valid = true
					break
				}
			}
			if valid == true {
				break
			}
		}

		if valid == false {
			fmt.Println("result:")
			fmt.Println(input[i])
			break
		}
	}
}

func extractInput() []int {
	fileBytes, err := ioutil.ReadFile("input/day9")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(string(fileBytes), "\n")

	var input []int
	for _, line := range lines {
		input = append(input, atoi(line))

	}

	return input
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return i
}
