package main

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func parserDayFour(inputConent string) ([][]string, error) {
	lines := strings.Split(inputConent, "\n")

	mainArray := [][]string{}
	for _, line := range lines {
		charSplit := strings.Split(line, "")
		mainArray = append(mainArray, charSplit)
	}

	return mainArray, nil
}

type Coordinates struct {
	x int
	y int
}

func starSearchCount(wordSearch [][]string, word string, startPoint Coordinates) (int, error) {

	// Get limits
	searchLen := len(word)
	rowLen := len(wordSearch)
	colLen := len(wordSearch[0])

	comboOptions := []Coordinates{
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: 1, y: 1},
		{x: -1, y: 0},
		{x: 0, y: -1},
		{x: -1, y: -1},
		{x: 1, y: -1},
		{x: -1, y: 1},
	}

	// Check if the word can fit in the grid
	if (startPoint.x+searchLen) > rowLen && (startPoint.x-searchLen) < 0 {
		return 0, errors.New("Word cannot fit in the X grid")
	}

	if (startPoint.y+searchLen) > colLen && (startPoint.y-searchLen) < 0 {
		return 0, errors.New("Word cannot fit in the Y grid")
	}

	searchCount := 0
	for _, comboValue := range comboOptions {

		xEnd := startPoint.x + comboValue.x
		yEnd := startPoint.y + comboValue.y

		searchCoords := []Coordinates{}
		for i := 0; i < searchLen; i++ {
			if comboValue.x > 0 {
				xEnd = startPoint.x + i
			}
			if comboValue.y > 0 {
				yEnd = startPoint.y + i
			}
			if comboValue.x < 0 {
				xEnd = startPoint.x - i
			}
			if comboValue.y < 0 {
				yEnd = startPoint.y - i
			}

			if xEnd >= rowLen || xEnd < 0 {
				continue
			}
			if yEnd >= colLen || yEnd < 0 {
				continue
			}

			tempCoords := Coordinates{x: xEnd, y: yEnd}
			searchCoords = append(searchCoords, tempCoords)
		}
		words := []string{}
		for _, searchCoord := range searchCoords {
			words = append(words, wordSearch[searchCoord.x][searchCoord.y])
		}

		if strings.Join(words, "") == word {
			searchCount++
		}

		slices.Reverse(words)
		if strings.Join(words, "") == word {
			searchCount++
		}

	}

	return searchCount, nil
}

func answerDayFour(wordSearch [][]string) int {

	count := 0
	for rIndex, row := range wordSearch {
		for cIndex, value := range row {
			if value == "X" {
				starCount, error := starSearchCount(wordSearch, "XMAS", Coordinates{x: rIndex, y: cIndex})
				if error != nil {
					continue
				}
				count += starCount
			}
		}
	}

	return count
}

func answerDayFourP2(wordSearch [][]string) int {

	foundPairs := []string{}
	for rIndex, row := range wordSearch {
		for cIndex, value := range row {
			if value == "A" {
				pairFound, error := searchMasX(wordSearch, Coordinates{x: rIndex, y: cIndex})
				if error != nil {
					continue
				}
				if len(pairFound) > 0 {
					fmt.Println("Pairs found: ", pairFound)
					foundPairs = append(foundPairs, pairFound...)
				}
			}
		}
	}

	// Distinct count of pairs
	fmt.Println("Number of Pairs found: ", len(foundPairs))
	distinctPairs := []string{}
	for _, pair := range foundPairs {
		if !slices.Contains(distinctPairs, pair) {
			distinctPairs = append(distinctPairs, pair)
		}
	}

	return len(distinctPairs)
}

func searchMasX(wordSearch [][]string, startPoint Coordinates) ([]string, error) {

	// Get limits
	targetWord := "MAS"
	searchLen := len(targetWord)
	rowLen := len(wordSearch)
	colLen := len(wordSearch[0])

	rightLeaning := []Coordinates{
		{x: 1, y: 1},   // top right
		{x: -1, y: -1}, // bottom left
	}

	leftLeaning := []Coordinates{
		{x: -1, y: 1}, // top left
		{x: 1, y: -1}, // bottom right
	}

	// Check if the word can fit in the grid
	if (startPoint.x+searchLen) > rowLen && (startPoint.x-searchLen) < 0 {
		return []string{""}, errors.New("Word cannot fit in the X grid")
	}

	if (startPoint.y+searchLen) > colLen && (startPoint.y-searchLen) < 0 {
		return []string{""}, errors.New("Word cannot fit in the Y grid")
	}

	// Right diagonal check
	rightSlice := []Coordinates{}
	for _, rightCheck := range rightLeaning {
		xEnd := startPoint.x + rightCheck.x
		yEnd := startPoint.y + rightCheck.y

		if xEnd >= rowLen || xEnd < 0 {
			return []string{""}, errors.New("Word cannot fit in the X grid")
		}
		if yEnd >= colLen || yEnd < 0 {
			return []string{""}, errors.New("Word cannot fit in the Y grid")
		}
		tempCoords := Coordinates{x: xEnd, y: yEnd}

		rightSlice = append(rightSlice, tempCoords)
	}
	rightSliceValues := []string{}
	for _, rightSliceValue := range rightSlice {
		rightSliceValues = append(rightSliceValues, wordSearch[rightSliceValue.x][rightSliceValue.y])
	}
	rightWord := rightSliceValues[0] + "A" + rightSliceValues[1]
	rightWordReverse := rightSliceValues[1] + "A" + rightSliceValues[0]

	// Left diagonal check
	leftSlice := []Coordinates{}
	for _, leftCheck := range leftLeaning {
		xEnd := startPoint.x + leftCheck.x
		yEnd := startPoint.y + leftCheck.y

		if xEnd >= rowLen || xEnd < 0 {
			return []string{""}, errors.New("Word cannot fit in the X grid")
		}
		if yEnd >= colLen || yEnd < 0 {
			return []string{""}, errors.New("Word cannot fit in the Y grid")
		}
		tempCoords := Coordinates{x: xEnd, y: yEnd}

		leftSlice = append(rightSlice, tempCoords)
	}
	leftSliceValues := []string{}
	for _, leftSliceValue := range leftSlice {
		leftSliceValues = append(leftSliceValues, wordSearch[leftSliceValue.x][leftSliceValue.y])
	}
	leftWord := leftSliceValues[0] + "A" + leftSliceValues[1]
	leftWordReverse := leftSliceValues[1] + "A" + leftSliceValues[0]

	// Return the coordinates if the word is found
	returningCoordString := []string{}

	if rightWord == "MAS" || rightWordReverse == "MAS" {
		tempValue := strconv.Itoa(rightSlice[0].x) + "," + strconv.Itoa(rightSlice[0].y)
		tempValue += "-" + strconv.Itoa(startPoint.x) + "," + strconv.Itoa(startPoint.y)
		tempValue += "-" + strconv.Itoa(rightSlice[1].x) + "," + strconv.Itoa(rightSlice[1].y)
		returningCoordString = append(returningCoordString, tempValue)

	}

	if leftWord == "MAS" || leftWordReverse == "MAS" {
		tempValue := strconv.Itoa(leftSlice[0].x) + "," + strconv.Itoa(leftSlice[0].y)
		tempValue += "-" + strconv.Itoa(startPoint.x) + "," + strconv.Itoa(startPoint.y)
		tempValue += "-" + strconv.Itoa(leftSlice[1].x) + "," + strconv.Itoa(leftSlice[1].y)
		returningCoordString = append(returningCoordString, tempValue)
	}
	return returningCoordString, nil
}

func foundSolutionOnlineDay4P2(grid [][]string) int {

	rows := len(grid)
	cols := len(grid[0])

	var count int
	var matchGrids = [][][]string{
		{{"M", "M"}, {"S", "S"}},
		{{"S", "S"}, {"M", "M"}},
		{{"S", "M"}, {"S", "M"}},
		{{"M", "S"}, {"M", "S"}}}

	for horiz := 0; horiz < cols; horiz++ {
		for vert := 0; vert < rows; vert++ {
			// boundary dodges
			x := horiz + 1
			y := vert + 1
			if y+1 < cols && x+1 < rows && grid[x][y] == "A" {
				tl := grid[x-1][y-1]
				tr := grid[x+1][y-1]
				bl := grid[x-1][y+1]
				br := grid[x+1][y+1]
				cmpGrid := [][]string{{tl, tr}, {bl, br}}
				for _, mg := range matchGrids {
					if reflect.DeepEqual(mg, cmpGrid) {
						count += 1
					}
				}

			}
		}
	}
	return count
}
