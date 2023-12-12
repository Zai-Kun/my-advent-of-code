package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var inputFileName = "input.txt"
var textualNumbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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

func getMinAndMaxIndex(m map[int]string) map[string]int{
	var min, max int
	for k := range m {
		min = k
		max = k
		break
	}

	for k := range m {
		if k < min {
			min = k
		}
		if k > max {
			max = k
		}
	}

	return map[string]int{"min": min, "max": max}
}

func firstIndexOf(text, subText string) int {
	return strings.Index(text, subText)
}

func lastIndexOf(text, subText string) int {
	for i := len(text) - len(subText); i >= 0; i-- {
		if strings.HasPrefix(text[i:], subText) {
			return i
		}
	}
	return -1
}

func extractNumbersFromText(text string) map[int]string {
	foundNumbers := map[int]string{}
	for i, char := range text {
		_, err := strconv.Atoi(string(char))
		if err == nil {
			foundNumbers[i] = string(char)
		}
	}

	// to find textual numbers
    for i := 0; i < len(textualNumbers); i++ {
    	firstOccurrenceIndex := firstIndexOf(text, textualNumbers[i])
    	if firstOccurrenceIndex == -1 {continue}
    	foundNumbers[firstOccurrenceIndex] = strconv.Itoa(i+1)

    	lastOccurrenceIndex := lastIndexOf(text, textualNumbers[i])
    	if lastOccurrenceIndex != firstOccurrenceIndex {
    		foundNumbers[lastOccurrenceIndex] = strconv.Itoa(i+1)
    	}

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

		maxAndMinIndex := getMinAndMaxIndex(extractedNumbers)
		firstNumber := extractedNumbers[maxAndMinIndex["min"]]
		lastNumber := extractedNumbers[maxAndMinIndex["max"]]

		finalNumber, err := strconv.Atoi(firstNumber + lastNumber)
		if err != nil {
			fmt.Println(err)
			return
		}

		totalSum += finalNumber
	}

	fmt.Println(totalSum)
}
