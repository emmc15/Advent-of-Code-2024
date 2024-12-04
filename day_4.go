package main

import (
	"errors"
	"slices"
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
