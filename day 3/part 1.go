package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const inputFile = "input.txt"

type Number struct{
	Value      int
	StartIndex int
	EndIndex   int
}

type Line struct{
	Numbers []*Number
	Symbols []int
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	totalSum := 0
	var prevLine *Line

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		processedLine, err := processLine(line)
		if err != nil{
			panic(err)
		}
		for _, number := range processedLine.Numbers{
			for _, symbolIndex := range processedLine.Symbols{
				if adjacentToSymbol(number, symbolIndex){
					totalSum += number.Value 
				}
			}
		}

		if prevLine != nil{
			for _, number := range prevLine.Numbers{
				for _, symbolIndex := range processedLine.Symbols{
					if adjacentToSymbol(number, symbolIndex){
						totalSum += number.Value 
					}
				}
			}
			for _, number := range processedLine.Numbers{
				for _, symbolIndex := range prevLine.Symbols{
					if adjacentToSymbol(number, symbolIndex){
						totalSum += number.Value 
					}
				}
			}
		}

		prevLine = processedLine
		
	}
	fmt.Println(totalSum)
}

func processLine(line string)(*Line, error){
	foundSymbols := []int{}
	foundNumbers := []*Number{}

	tmpBuffer := ""
	StartIndex := -1
	EndIndex := -1
	
	for i, char := range line{
		charStr := string(char)
		isDigit := unicode.IsDigit(char)
		if charStr == "." && tmpBuffer == ""{
			continue
		}

		EndIndex = i-1
		if isDigit{
			if tmpBuffer == "" {
				StartIndex = i
			}

			tmpBuffer += charStr
			if len(line)-1 != i{
				continue
			}
			EndIndex = i
		}

		if charStr != "." && tmpBuffer == ""{
			foundSymbols = append(foundSymbols, i)
			continue
		}

		number := &Number{}

		value, err := strconv.Atoi(tmpBuffer)
		if err != nil {return nil, err}

		number.Value = value
		number.StartIndex = StartIndex
		number.EndIndex = EndIndex
		foundNumbers = append(foundNumbers, number)

		tmpBuffer = ""
		if charStr != "." && !isDigit{
			foundSymbols = append(foundSymbols, i)
		}
	}

	processedLine := &Line{}
	processedLine.Numbers = foundNumbers
	processedLine.Symbols = foundSymbols

	return processedLine, nil
}

func adjacentToSymbol(number *Number, symbolIndex int)bool{
	start := number.StartIndex
	if number.StartIndex > 0 {
		start--
	}

	return symbolIndex>=start && symbolIndex<=number.EndIndex+1
}
