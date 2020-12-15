package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var invalid int
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
			invalid = input[i]
			break
		}
	}
	fmt.Println("result part 1:")
	fmt.Println(invalid)

	for i := 0; i < len(input); i++ {
		candidates := []int{input[i]}
		candidateKeyOffset := 1
		sum := input[i]

		for sum < invalid {
			candidateKey := i + candidateKeyOffset
			sum += input[candidateKey]
			candidates = append(candidates, input[candidateKey])
			candidateKeyOffset++
		}

		if sum == invalid {
			sort.Ints(candidates)
			fmt.Println("result part 2:")
			fmt.Println(candidates[0] + candidates[len(candidates)-1])
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
