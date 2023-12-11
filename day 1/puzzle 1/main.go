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

func extractNumbersFromText(text string) []string {
	foundNumbers := []string{}
	for _, char := range text {
		_, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}

		foundNumbers = append(foundNumbers, string(char))
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

		firstNumber := extractedNumbers[0]
		lastNumber := extractedNumbers[len(extractedNumbers)-1]

		finalNumber, err := strconv.Atoi(firstNumber + lastNumber)
		if err != nil {
			fmt.Println(err)
			return
		}

		totalSum += finalNumber
	}

	fmt.Println(totalSum)
}
