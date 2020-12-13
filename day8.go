package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	instructions := extractInput()

	for key, instruction := range instructions {
		testedInstructions := make([]string, len(instructions))
		copy(testedInstructions, instructions)
		action := strings.Split(instruction, " ")
		switch action[0] {
		case "nop":
			testedInstructions[key] = "jmp " + action[1]
		case "acc":
			continue
		case "jmp":
			testedInstructions[key] = "nop " + action[1]
		default:
			fmt.Printf("meh")
		}

		executeInstructions(testedInstructions)
	}
}

func executeInstructions(instructions []string) {
	var alreadyExecutedKey []int
	key := 0
	acc := 0
	for {
		instruction := strings.Split(instructions[key], " ")
		alreadyExecutedKey = append(alreadyExecutedKey, key)
		switch instruction[0] {
		case "nop":
			key++
		case "acc":
			acc += atoi(instruction[1])
			key++
		case "jmp":
			key += atoi(instruction[1])
		default:
			fmt.Printf("meh")
		}

		if contains(alreadyExecutedKey, key) {
			break
		}

		if key >= len(instructions) {
			fmt.Println("Finish ! Value in the accumulator: " + strconv.Itoa(acc))
			os.Exit(0)
			break
		}
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return i
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func extractInput() []string {
	fileBytes, err := ioutil.ReadFile("input/day8")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	instructions := strings.Split(string(fileBytes), "\n")

	return instructions
}
