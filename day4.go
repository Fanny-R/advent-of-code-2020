package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passports := extractPassports()
	fmt.Println(len(passports))
	count := 0
	for _, passport := range passports {
		if (len(passport)) < 7 {

			continue
		}

		byr, ok := passport["byr"]
		if !(ok && isValueValid(byr, 1920, 2002)) {

			continue
		}

		iyr, ok := passport["iyr"]
		if !(ok && isValueValid(iyr, 2010, 2020)) {

			continue
		}

		eyr, ok := passport["eyr"]
		if !(ok && isValueValid(eyr, 2020, 2030)) {

			continue
		}

		hgt, ok := passport["hgt"]
		if !ok {

			continue
		}
		last2 := hgt[len(hgt)-2:]
		size := hgt[0 : len(hgt)-2]
		if last2 == "cm" {
			if !isValueValid(size, 150, 193) {
				continue
			}
		} else if last2 == "in" {
			if !isValueValid(size, 59, 76) {
				continue
			}
		} else {
			continue
		}

		hcl, ok := passport["hcl"]
		if !ok {

			continue
		}
		matched, err := regexp.Match(`#([0-9]|[a-f]){6}`, []byte(hcl))

		if !matched || err != nil {
			continue
		}

		ecl, ok := passport["ecl"]
		if !(ok && isValidEyeColor(ecl)) {
			continue
		}

		pid, ok := passport["pid"]

		if !(ok && len(pid) == 9) {

			continue
		}

		matchedPid, err := regexp.Match(`[0-9]{9}`, []byte(pid))

		if !matchedPid || err != nil {
			continue
		}

		count++
	}

	fmt.Println(count)
}

func isValidEyeColor(color string) bool {
	validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, validColor := range validColors {
		if color == validColor {
			return true
		}
	}
	return false
}

func isValueValid(val string, min int, max int) bool {
	i, err := strconv.Atoi(val)
	if err != nil {
		return false
	}

	if i < min || i > max {
		return false
	}
	return true
}

func extractPassports() []map[string]string {
	fileBytes, err := ioutil.ReadFile("input/day4")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(string(fileBytes), "\n\n")
	passports := make([]map[string]string, 0)

	for _, line := range lines {
		var passportLine []string
		halfSplittedLine := strings.Split(line, "\n")
		for _, halfLine := range halfSplittedLine {
			passportLine = append(passportLine, strings.Split(halfLine, " ")...)
		}

		passport := make(map[string]string)
		for _, element := range passportLine {
			field := strings.Split(element, ":")
			passport[field[0]] = field[1]
		}

		passports = append(passports, passport)
	}

	return passports
}
