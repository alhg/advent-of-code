package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Password struct {
	minNumChar   int
	maxNumChar   int
	requiredChar string
	password     string
}

func getLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	inputs := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}

	return inputs
}

func parseInput(inputs []string) []Password {
	passwords := []Password{}
	for _, input := range inputs {
		var minChar int
		var maxChar int
		var requiredChar string
		var aPassword string
		_, err := fmt.Sscanf(input, "%d-%d %s %s", &minChar, &maxChar, &requiredChar, &aPassword)
		requiredChar = strings.ReplaceAll(requiredChar, ":", "")
		if err != nil {
			panic(err)
		}
		inputPassword := Password{minNumChar: minChar, maxNumChar: maxChar, requiredChar: requiredChar, password: aPassword}
		passwords = append(passwords, inputPassword)
	}
	return passwords
}

func getNumValidPasswordsP1(passwords []Password) int {
	numValidPasswords := 0
	for _, password := range passwords {
		numRequiredChar := strings.Count(password.password, password.requiredChar)
		if numRequiredChar >= password.minNumChar && numRequiredChar <= password.maxNumChar {
			numValidPasswords++
		}
	}
	return numValidPasswords
}

func getNumValidPasswordsP2(passwords []Password) int {
	numValidPasswords := 0
	for _, password := range passwords {
		firstIndex := password.minNumChar - 1
		secondIndex := password.maxNumChar - 1

		if len(password.password) <= firstIndex || len(password.password) <= secondIndex {
			continue
		}

		// Exclusive OR comparison
		if (string(password.password[firstIndex]) == password.requiredChar) !=
			(string(password.password[secondIndex]) == password.requiredChar) {
			numValidPasswords++
		}
	}
	return numValidPasswords
}

func main() {
	inputs := getLines("../input.txt")
	passwords := parseInput(inputs)
	numValidPasswordsP1 := getNumValidPasswordsP1(passwords)
	fmt.Println("Part One: Number of Valid Passwords:", numValidPasswordsP1)
	numValidPasswordsP2 := getNumValidPasswordsP2(passwords)
	fmt.Println("Part Two: Number of Valid Passwords:", numValidPasswordsP2)
}
