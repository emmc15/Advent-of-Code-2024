package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func parserDayOne(inputConent string) ([]int, []int, error) {
	lines := strings.Split(inputConent, "\n")

	// turn lines into two lists
	leftList := []int{}
	rightList := []int{}
	for _, line := range lines {
		// do something with the line
		stringSplit := strings.Split(line, "   ")
		leftVal := stringSplit[0]
		rightVal := stringSplit[1]

		leftValInt, err := strconv.Atoi(leftVal)
		if err != nil {
			return nil, nil, err
		}

		rightValInt, err := strconv.Atoi(rightVal)

		if err != nil {
			return nil, nil, err
		}

		leftList = append(leftList, leftValInt)
		rightList = append(rightList, rightValInt)
	}

	return leftList, rightList, nil

}

func answerDayOne(leftList []int, rightList []int) int {
	// sort the list of ints to min first
	// iterate through the list of ints

	slices.Sort(leftList)
	slices.Sort(rightList)
	// distances := []int{}
	sum := 0
	for index, _ := range leftList {
		leftVal := leftList[index]
		rightVal := rightList[index]

		distance := leftVal - rightVal
		distance = int(math.Abs(float64(distance)))
		// distances = append(distances, distance)
		sum += distance

		fmt.Println("Distance: ", distance, "Sum: ", sum)

	}

	return sum
}

func answerDayOneP2(leftList []int, rightList []int) int {
	// sort the list of ints to min first
	// iterate through the list of ints

	slices.Sort(leftList)
	slices.Sort(rightList)

	similarity := 0

	for _, leftVal := range leftList {
		leftValCount := 0

		for _, rightVal := range rightList {
			if leftVal == rightVal {
				leftValCount++
			}
		}

		similarity += (leftValCount * leftVal)
	}

	return similarity
}
