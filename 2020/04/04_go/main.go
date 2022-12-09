package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func getPassportStrs(filename string) []string {
	content, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	text := string(content)
	passportStrs := strings.Split(text, "\n\n")
	return passportStrs
}

func getValidPassportsP1(passportStrs []string) int {
	validKeys := []string{
		"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
	}
	numValidPassports := 0
	for _, passport := range passportStrs {
		keyValuePairs := strings.Fields(passport)
		passportMap := make(map[string]string)
		for _, keyValuePair := range keyValuePairs {
			keyValueArr := strings.Split(keyValuePair, ":")
			passportMap[keyValueArr[0]] = keyValueArr[1]
		}
		isValidPassportP1 := true
		for _, validKey := range validKeys {
			_, isInMap := passportMap[validKey]
			if !isInMap {
				isValidPassportP1 = false
				break
			}
		}
		if isValidPassportP1 {
			numValidPassports++
		}
	}
	return numValidPassports
}

func getValidPassportsP2(passportStrs []string) int {
	numValidPassports := 0
	for _, passport := range passportStrs {
		keyValuePairs := strings.Fields(passport)
		passportMap := make(map[string]string)
		for _, keyValuePair := range keyValuePairs {
			keyValueArr := strings.Split(keyValuePair, ":")
			passportMap[keyValueArr[0]] = keyValueArr[1]
		}
		if isValidPassportP2(passportMap) {
			numValidPassports++
		}
	}
	return numValidPassports
}

func isValidPassportP2(passport map[string]string) bool {
	validKeysAndValidatorMap := map[string]func(string) bool{
		"byr": isValidBirthYear,
		"iyr": isValidIssueYear,
		"eyr": isValidExpirationYear,
		"hgt": isValidHeight,
		"hcl": isValidHairColor,
		"ecl": isValidEyeColor,
		"pid": isValidPassportID,
	}
	for validKey, validator := range validKeysAndValidatorMap {
		value, isInMap := passport[validKey]
		if !isInMap || !validator(value) {
			// panic(fmt.Errorf(""))
			return false
		} else if validator(value) {
			fmt.Printf("%q %q\n", validKey, value)
		}
	}
	return true
}

func isValidBirthYear(birthYear string) bool {
	if len(birthYear) != 4 {
		return false
	}

	year, err := strconv.Atoi(birthYear)
	if err != nil {
		return false
	}

	if year < 1920 || year > 2002 {
		return false
	}
	return true
}

func isValidIssueYear(issueYear string) bool {
	if len(issueYear) != 4 {
		return false
	}

	year, err := strconv.Atoi(issueYear)
	if err != nil {
		return false
	}

	if year < 2010 || year > 2020 {
		return false
	}
	return true
}

func isValidExpirationYear(expirationYear string) bool {
	if len(expirationYear) != 4 {
		return false
	}

	year, err := strconv.Atoi(expirationYear)
	if err != nil {
		return false
	}

	if year < 2020 || year > 2030 {
		return false
	}
	return true
}

func isValidHeight(heightStr string) bool {
	var height int
	var unit string

	n, err := fmt.Sscanf(heightStr, "%d%s", &height, &unit)
	if err != nil || n != 2 {
		return false
	}

	if unit != "cm" && unit != "in" {
		return false
	}
	if unit == "cm" && (height < 150 || height > 193) {
		return false
	}
	if unit == "in" && (height < 59 || height > 76) {
		return false
	}
	return true
}

func isValidHairColor(hairColorStr string) bool {
	if len(hairColorStr) != 7 {
		return false
	}
	if string(hairColorStr[0]) != "#" {
		return false
	}
	for _, digitOrChar := range hairColorStr[1:] {
		if !unicode.IsDigit(digitOrChar) && (digitOrChar < 'a' || digitOrChar > 'f') {
			return false
		}
	}
	return true
}

func isValidEyeColor(eyeColourStr string) bool {
	validEyeColors := []string{
		"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
	}
	for _, validEyeColor := range validEyeColors {
		if eyeColourStr == validEyeColor {
			return true
		}
	}
	return false
}

func isValidPassportID(passportID string) bool {
	if len(passportID) != 9 {
		return false
	}
	for _, digit := range passportID {
		if !unicode.IsDigit(digit) {
			return false
		}
	}
	return true
}

func main() {
	passportStrs := getPassportStrs("../input.txt")
	numValidPassportsP1 := getValidPassportsP1(passportStrs)
	numValidPassportsP2 := getValidPassportsP2(passportStrs)

	fmt.Println("Day 04 - Part 1 - # of Valid Passports:", numValidPassportsP1)
	fmt.Println("Day 04 - Part 2 - # of Valid Passports:", numValidPassportsP2)
}
