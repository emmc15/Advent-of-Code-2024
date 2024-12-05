// simple cli tool that runs a hello to console

package main

import (
	"fmt"
	"os"
)

func readProblemTxt(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// read the file
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)

	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func solutionDayOne() {
	problemTxt, err := readProblemTxt("./puzzle_sets/day_1a.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	leftArray, rightArray, err := parserDayOne(problemTxt)
	if err != nil {
		fmt.Println("Error parsing file: ", err)
		return
	}

	answer := answerDayOne(leftArray, rightArray)
	fmt.Println("Day1, P1 Answer: ", answer)

	answerP2 := answerDayOneP2(leftArray, rightArray)
	fmt.Println("Day1, P2 Answer: ", answerP2)

}

func solutionDayTwo() {
	problemTxt, err := readProblemTxt("./puzzle_sets/day_2a.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	matrix, err := parserDayTwo(problemTxt)
	if err != nil {
		fmt.Println("Error parsing file: ", err)
		return
	}

	answer := answerDayTwo(matrix)
	fmt.Println("Day2, P1 Answer: ", answer)

	answerP2 := answerDayTwoP2(matrix)
	fmt.Println("Day2, P2 Answer: ", answerP2)
}

func solutionDayThree() {
	problemTxt, err := readProblemTxt("./puzzle_sets/day_3a.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	matrix, err := parserDayThree(problemTxt)
	if err != nil {
		fmt.Println("Error parsing file: ", err)
		return
	}

	answer := answerDayThree(matrix)
	fmt.Println("Day3, P1 Answer: ", answer)

	matrix2, err := parserDayThreeP2(problemTxt)
	if err != nil {
		fmt.Println("Error parsing file: ", err)
		return
	}

	answer2 := answerDayThree(matrix2)
	fmt.Println("Day3, P2 Answer: ", answer2)

}

func solutionDayFour() {
	problemTxt, err := readProblemTxt("./puzzle_sets/day_4.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	matrix, err := parserDayFour(problemTxt)
	if err != nil {
		fmt.Println("Error parsing file: ", err)
		return
	}

	answer := answerDayFour(matrix)
	fmt.Println("Day4, P1 Answer: ", answer)

	// answerP2 := answerDayFourP2(matrix)
	answerP2 := foundSolutionOnlineDay4P2(matrix)
	fmt.Println("Day4, P2 Answer: ", answerP2)

}

func main() {
	solutionDayOne()
	fmt.Println()
	solutionDayTwo()
	fmt.Println()
	solutionDayThree()
	fmt.Println()
	solutionDayFour()
}
