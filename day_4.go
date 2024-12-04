package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type GridValue struct {
	value  string
	rIndex int
	cIndex int
}

func parserDayFour(inputConent string) ([][]string, error) {
	lines := strings.Split(inputConent, "\n")

	mainArray := [][]string{}
	for _, line := range lines {
		charSplit := strings.Split(line, "")
		mainArray = append(mainArray, charSplit)
	}

	return mainArray, nil
}

func checkXmasStringIsValid(inputChar string, xmasArray []string) bool {
	xmas := []string{"X", "M", "A", "S"}

	if !slices.Contains(xmas, inputChar) {
		return false
	}

	if len(xmasArray) == 0 {
		return false
	}

	inputCharIndex := slices.Index(xmas, inputChar)

	// if xmasArrayLength == 1 {
	// 	if inputCharIndex == 0 {
	// 		return true
	// 	}
	// 	return false
	// }

	for _, xChar := range xmasArray {
		xCharIndex := slices.Index(xmas, xChar)

		if xCharIndex == -1 {
			return false
		}

		if (inputCharIndex - 1) == xCharIndex {
			return true
		}

		if (inputCharIndex + 1) == xCharIndex {
			return true
		}
	}

	return true
}

func getXmasMapping(inputChar string, xmasArray []GridValue) (GridValue, error) {
	xmas := []string{"X", "M", "A", "S"}

	if !slices.Contains(xmas, inputChar) {
		return GridValue{}, errors.New("Invalid input character")
	}

	if len(xmasArray) == 0 {
		return GridValue{}, errors.New("Invalid xmas array")
	}

	inputCharIndex := slices.Index(xmas, inputChar)

	for xmasArrayIndex, xChar := range xmasArray {
		xCharIndex := slices.Index(xmas, xChar)

		if xCharIndex == -1 {
			return "", -1
		}

		if (inputCharIndex + 1) == xCharIndex {
			return xChar, xmasArrayIndex
		}

		if (inputCharIndex - 1) == xCharIndex {
			return xChar, xmasArrayIndex
		}
	}

	return "", -1
}

func getGridAdjescentValues(grid [][]string, rIndex int, cIndex int) []GridValue {
	// get the values of the grid
	// get the values that are adjescent to the current value
	// return the values
	// grid[rIndex][cIndex]
	// grid[rIndex][cIndex+1]
	// grid[rIndex][cIndex-1]
	// grid[rIndex+1][cIndex]
	// grid[rIndex-1][cIndex]
	// grid[rIndex+1][cIndex+1]
	// grid[rIndex+1][cIndex-1]
	// grid[rIndex-1][cIndex+1]
	// grid[rIndex-1][cIndex-1]

	// get the values of the grid
	gridValues := []string{}
	for _, row := range grid {
		for _, value := range row {
			gridValues = append(gridValues, value)
		}
	}

	// get the values that are adjescent to the current value
	adjescentValues := []GridValue{}
	if cIndex < len(grid[rIndex])-1 {
		foundValue := GridValue{value: grid[rIndex][cIndex+1], rIndex: rIndex, cIndex: cIndex + 1}
		adjescentValues = append(adjescentValues, foundValue)
	}
	if cIndex > 0 {
		foundValue := GridValue{value: grid[rIndex][cIndex-1], rIndex: rIndex, cIndex: cIndex - 1}
		adjescentValues = append(adjescentValues, foundValue)
	}
	if rIndex < len(grid)-1 {
		foundValue := GridValue{value: grid[rIndex+1][cIndex], rIndex: rIndex + 1, cIndex: cIndex}
		adjescentValues = append(adjescentValues, foundValue)
	}
	if rIndex > 0 {
		foundValue := GridValue{value: grid[rIndex-1][cIndex], rIndex: rIndex - 1, cIndex: cIndex}
		adjescentValues = append(adjescentValues, foundValue)
	}
	if rIndex < len(grid)-1 && cIndex < len(grid[rIndex])-1 {
		foundValue := GridValue{value: grid[rIndex+1][cIndex+1], rIndex: rIndex + 1, cIndex: cIndex + 1}
		adjescentValues = append(adjescentValues, foundValue)
	}
	if rIndex < len(grid)-1 && cIndex > 0 {
		foundValue := GridValue{value: grid[rIndex+1][cIndex-1], rIndex: rIndex + 1, cIndex: cIndex - 1}
		adjescentValues = append(adjescentValues, foundValue)
	}
	if rIndex > 0 && cIndex < len(grid[rIndex])-1 {
		foundValue := GridValue{value: grid[rIndex-1][cIndex+1], rIndex: rIndex - 1, cIndex: cIndex + 1}
		adjescentValues = append(adjescentValues, foundValue)
	}
	if rIndex > 0 && cIndex > 0 {
		foundValue := GridValue{value: grid[rIndex-1][cIndex-1], rIndex: rIndex - 1, cIndex: cIndex - 1}
		adjescentValues = append(adjescentValues, foundValue)
	}

	// return the values
	return adjescentValues

}

func answerDayFour(wordSearch [][]string) int {

	filteredWordSearch := [][]string{}
	for rIndex, row := range wordSearch {
		filteredRow := []string{}
		for cIndex, value := range row {

			placeHolder := "."
			adjescentValues := getGridAdjescentValues(wordSearch, rIndex, cIndex)

			listValues := []string{}
			for _, adjescentValue := range adjescentValues {
				fmt.Println(adjescentValue)

				listValues = append(listValues, adjescentValue.value)
			}
			matchedValue, adjescentIndex := getXmasMapping(value, listValues)
			fmt.Println(matchedValue, adjescentIndex)

			if matchedValue == "" {
				filteredRow = append(filteredRow, value)
				continue
			}

			filteredRow = append(filteredRow, placeHolder)

		}
		filteredWordSearch = append(filteredWordSearch, filteredRow)
	}

	fmt.Println(filteredWordSearch)
	return 0

}
