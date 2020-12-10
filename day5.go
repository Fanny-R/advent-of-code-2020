package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	boardingPasses := extractInput()
	highestSeatId := 0
	var seatIdsList []int
	for _, boardingPass := range boardingPasses {
		row := findPosition([]int{0, 127}, boardingPass[0:7])
		column := findPosition([]int{0, 7}, boardingPass[7:10])

		seatId := (row * 8) + column

		seatIdsList = append(seatIdsList, seatId)

		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}

	fmt.Println("Result")
	fmt.Println(highestSeatId)

	sort.Ints(seatIdsList)

	var mySeat int
	for key, seatId := range seatIdsList {
		if seatIdsList[key+1] != seatId+1 {
			mySeat = seatId + 1
			break
		}
	}

	fmt.Println("My seat")
	fmt.Println(mySeat)
}

func findPosition(rangeEl []int, instructions string) int {
	for _, instruction := range instructions {
		instruction := string(instruction)
		if instruction == "F" || instruction == "L" {
			rangeEl[1] = rangeEl[0] + ((rangeEl[1] - rangeEl[0]) / 2)
		} else if instruction == "B" || instruction == "R" {
			rangeEl[0] = rangeEl[0] + ((rangeEl[1]-rangeEl[0])/2 + 1)
		}
	}

	if rangeEl[0] != rangeEl[1] {
		fmt.Println("Oh no")
		os.Exit(1)
	}

	return rangeEl[0]
}

func extractInput() []string {
	fileBytes, err := ioutil.ReadFile("input/day5")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(fileBytes), "\n")
}
