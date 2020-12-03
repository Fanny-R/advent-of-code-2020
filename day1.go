package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	// input := []int{1721, 979, 366, 299, 675, 1456}
	input := extractInput()

	for i := 0; i < len(input); i++ {
		for j := 1; j < len(input); j++ {
			for k := 2; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					fmt.Println("YAY")
					fmt.Println(input[i])
					fmt.Println(input[j])
					fmt.Println(input[k])
					fmt.Println(input[i] * input[j] * input[k])
					return
				}
			}
		}
	}
}

func extractInput() []int {
	fileBytes, err := ioutil.ReadFile("input/day1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")

	var input = []int{}

	for _, i := range sliceData {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		input = append(input, j)
	}

	return input
}
