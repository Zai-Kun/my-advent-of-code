package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var inputFileName = "input.txt"

func readInputFile() (string, error) {
	file, err := os.Open(inputFileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(fileBytes), nil

}

func extractNumbersFromText(text string) []int {
	foundNumbers := []int{}
	for _, char := range text {
		number, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}

		foundNumbers = append(foundNumbers, number)
	}
	return foundNumbers
}

func main() {
	fileString, err := readInputFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	totalSum := 0
	for _, line := range strings.Split(fileString, "\n") {
		extractedNumbers := extractNumbersFromText(line)
		if len(extractedNumbers) == 0 {
			continue
		}

		firstNumber := strconv.Itoa(extractedNumbers[0])
		lastNumber := strconv.Itoa(extractedNumbers[len(extractedNumbers)-1])

		finalNumber, err := strconv.Atoi(firstNumber + lastNumber)
		if err != nil {
			fmt.Println(err)
			return
		}

		totalSum += finalNumber
	}

	fmt.Println(totalSum)
}
