package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func numOfTrees(inputs []string, right, down int) int {
	lineSize := len(inputs[0])
	numLines := len(inputs)
	currentRow := 0
	currentCol := 0
	numOfTrees := 0
	for {
		currentCol += right
		currentRow += down
		if currentRow >= numLines {
			break
		}
		if currentCol >= lineSize {
			currentCol = currentCol - lineSize
		}
		if string(inputs[currentRow][currentCol]) == "#" {
			numOfTrees++
		}
	}
	return numOfTrees
}

func partTwoSolution(inputs []string) int {
	type slope struct {
		right int
		down  int
	}
	slopes := []slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}
	product := 1
	for _, slope := range slopes {
		numOfTrees := numOfTrees(inputs, slope.right, slope.down)
		product *= numOfTrees
	}
	return product
}

func main() {
	inputs := getLines("../input.txt")
	fmt.Println("Day 03 - Part One - Num Of Trees:", numOfTrees(inputs, 3, 1))
	fmt.Println("Day 03 - Part Two - Num Of Trees:", partTwoSolution(inputs))
}
