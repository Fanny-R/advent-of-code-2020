package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// input := []string{"..##.......", "#...#...#..", ".#....#..#.", "..#.#...#.#", ".#...##..#.", "..#.##.....", ".#.#.#....#", ".#........#", "#.##...#...", "#...##....#", ".#..#...#.#"}
	input := extractInput()

	totalTrees := 1

	for p := 1; p <= 7; p += 2 {
		down := 1
		if p == 1 {
			down = 2
		}

		for d := 1; d <= down; d++ {
			currentPosition := p
			trees := 0
			for i := d; i < len(input); i += d {
				line := input[i]

				if string(line[currentPosition]) == "#" {
					trees++
				}

				currentPosition = currentPosition + p

				if currentPosition >= len(line) {
					currentPosition = currentPosition - len(line)
				}
			}

			fmt.Println(trees)
			totalTrees *= trees
		}
	}
	fmt.Println(totalTrees)
}

func extractInput() []string {
	fileBytes, err := ioutil.ReadFile("input/day3")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(fileBytes), "\n")
}
