package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	passports := extractPassports()
	fmt.Println(len(passports))
	count := 0
	for _, passport := range passports {
		if (len(passport)) >= 8 {
			count++
			continue
		}

		_, ok := passport["cid"]
		if ((len(passport)) == 7) && !(ok) {
			count++
			continue
		}

	}

	fmt.Println(count)
}

func extractPassports() []map[string]interface{} {
	fileBytes, err := ioutil.ReadFile("input/day4")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(string(fileBytes), "\n\n")
	passports := make([]map[string]interface{}, 0)

	for _, line := range lines {
		var passportLine []string
		halfSplittedLine := strings.Split(line, "\n")
		for _, halfLine := range halfSplittedLine {
			passportLine = append(passportLine, strings.Split(halfLine, " ")...)
		}

		passport := make(map[string]interface{})
		for _, element := range passportLine {
			field := strings.Split(element, ":")
			passport[field[0]] = field[1]
		}

		passports = append(passports, passport)
	}

	return passports
}
