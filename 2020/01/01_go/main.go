package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne(inputs []int) (int, error) {
	for _, n := range inputs {
		for _, n2 := range inputs {
			if n+n2 == 2020 {
				return n * n2, nil
			}
		}
	}
	return 0, fmt.Errorf("not found")
}

func partTwo(inputs []int) (int, error) {
	for _, n := range inputs {
		for _, n2 := range inputs {
			for _, n3 := range inputs {
				if n+n2+n3 == 2020 {
					return n * n2 * n3, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("not found")
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	inputs := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.Trim(line, "\n")
		num, err := strconv.Atoi(trimmedLine)
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, num)
	}

	product, err := partOne(inputs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One: %d\n", product)

	product, err = partTwo(inputs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part Two: %d\n", product)
}
