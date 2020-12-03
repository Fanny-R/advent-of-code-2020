package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//input := []string{"..##.......", "#...#...#..", ".#....#..#.", "..#.#...#.#", ".#...##..#.", "..#.##.....", ".#.#.#....#", ".#........#", "#.##...#...", "#...##....#", ".#..#...#.#"}
	input := extractInput()

	trees := 0
	currentPosition := 3
	for i, line := range input {
		if i == 0 {
			continue
		}

		if string(line[currentPosition]) == "#" {
			trees++
		}

		currentPosition = currentPosition + 3

		if currentPosition >= len(line) {
			currentPosition = currentPosition - len(line)
		}
	}

	fmt.Println(trees)
}

func extractInput() []string {
	fileBytes, err := ioutil.ReadFile("input/day3")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(fileBytes), "\n")
}
