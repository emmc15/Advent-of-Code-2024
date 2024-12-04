package main

import (
	"regexp"
	"strconv"
	"strings"
)

func parserDayThree(inputConent string) ([][]int, error) {
	r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	commands := r.FindAllString(inputConent, -1)

	mainArray := [][]int{}
	for _, command := range commands {
		command = strings.Replace(command, "mul(", "", -1)
		command = strings.Replace(command, ")", "", -1)
		numbers := strings.Split(command, ",")

		numbersInt := []int{}
		for _, number := range numbers {
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
			numbersInt = append(numbersInt, numberInt)
		}

		mainArray = append(mainArray, numbersInt)

	}

	return mainArray, nil

}

func parserDayThreeP2(inputConent string) ([][]int, error) {
	r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	commands := r.FindAllString(inputConent, -1)

	mainArray := [][]int{}
	multiFlag := true
	for _, command := range commands {
		if command == "do()" {
			multiFlag = true
			continue
		}
		if command == "don't()" {
			multiFlag = false
			continue
		}

		if multiFlag == false {
			continue
		}

		command = strings.Replace(command, "mul(", "", -1)
		command = strings.Replace(command, ")", "", -1)
		numbers := strings.Split(command, ",")

		numbersInt := []int{}
		for _, number := range numbers {
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
			numbersInt = append(numbersInt, numberInt)
		}

		mainArray = append(mainArray, numbersInt)

	}

	return mainArray, nil

}

func answerDayThree(matrix [][]int) int {
	// iterate through the matrix
	// perform the multiplication
	// sum the results
	sum := 0
	for _, row := range matrix {
		product := row[0] * row[1]
		sum += product
	}

	return sum
}
