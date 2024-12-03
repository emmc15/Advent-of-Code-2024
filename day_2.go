package main

import (
	"fmt"
	"strings"
)

func parserDayTwo(inputContent string) ([][]int, error) {
	lines := strings.Split(inputContent, "\n")

	// turn lines into two lists
	matrix := [][]int{}
	for _, line := range lines {
		// do something with the line
		stringSplit := strings.Split(line, " ")
		row := []int{}
		for _, val := range stringSplit {
			valInt := 0
			fmt.Sscanf(val, "%d", &valInt)
			row = append(row, valInt)
		}
		matrix = append(matrix, row)
	}

	return matrix, nil
}

func checkRowIsValid(row []int) bool {
	prevVal := row[0]
	direction := 0

	for index, val := range row {
		if index == 0 {
			continue
		}

		diff := val - prevVal

		// store the direction
		if index == 1 {
			direction = diff
		} else {
			// check direction
			if direction < 0 && diff > 0 {
				return false
			}
			if direction > 0 && diff < 0 {
				return false
			}
		}

		// check if the diff is at least 3
		if diff < 0 {
			diff = -diff
		}

		if diff > 3 || diff == 0 {
			return false
		}

		// reset the prevVal
		prevVal = val
	}
	return true
}

func checkRowIsValidWithTolerance(row []int, levelTollerance int) bool {
	prevVal := row[0]
	direction := 0

	toleranceCount := 0

	for index, val := range row {
		if index == 0 {
			continue
		}

		diff := val - prevVal
		if diff == 0 {
			toleranceCount++
			prevVal = val
			continue
		}

		// store the direction
		if index == 1 {
			direction = diff
		} else {
			// check direction
			if direction < 0 && diff > 0 {
				toleranceCount++
				continue
			}
			if direction > 0 && diff < 0 {
				toleranceCount++
				continue
			}
		}

		// check if the diff is at least 3
		if diff < 0 {
			diff = -diff
		}

		if diff > 3 || diff == 0 {
			toleranceCount++
			continue
		}

		// reset the prevVal
		prevVal = val
	}

	return toleranceCount <= levelTollerance
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func checkRowIsValidBruteForce(row []int) bool {

	if checkRowIsValid(row) {
		return true
	}

	// Remove one at a time and check if it is valid

	for index, _ := range row {
		newRow := []int{}
		//newRow = removeIndex(row, index)
		//fmt.Println("Old Row: ", row)
		//fmt.Println("New Row: ", newRow)

		startPart := row[:index]
		endPart := row[index+1:]
		for _, val := range startPart {
			newRow = append(newRow, val)
		}
		for _, val := range endPart {
			newRow = append(newRow, val)
		}

		if checkRowIsValid(newRow) {
			return true
		}
	}
	return false
}

func answerDayTwo(matrix [][]int) int {

	// look through each row to find it is a scale and diff is at least or at most 3

	safeRowCount := 0

	for _, row := range matrix {
		if checkRowIsValid(row) {
			safeRowCount++
		}

	}
	return safeRowCount
}

func answerDayTwoP2(matrix [][]int) int {

	// look through each row to find it is a scale and diff is at least or at most 3

	safeRowCount := 0

	for _, row := range matrix {
		if checkRowIsValidBruteForce(row) {
			safeRowCount++
		}

	}
	return safeRowCount
}
