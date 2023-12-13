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

type Symbol struct{
	Symbol string
	Index  int
}

type Line struct{
	LineNumber int
	Numbers    []*Number
	Symbols    []*Symbol
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	processedLines := []*Line{}

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		processedLine, err := processLine(line)
		processedLine.LineNumber = i
		if err != nil{panic(err)}

		processedLines = append(processedLines, processedLine)
		i++
	}

	fmt.Println(getTotalGearRatio(processedLines))
}

func getTotalGearRatio(lines []*Line) int {
	total := 0
	for _, line := range lines{
		for _, symbol := range line.Symbols{
			if symbol.Symbol != "*"{
				continue
			}

			start := 0
			end := line.LineNumber+1
			if line.LineNumber>0{start=line.LineNumber-1}
			if line.LineNumber+2<=len(lines){end=line.LineNumber+2}
			connected := []*Number{}
			for _, line := range lines[start:end]{
				for _, number := range line.Numbers{
					if adjacentToSymbol(number, symbol.Index){
						connected = append(connected, number)
					}
				}
			}
			if len(connected) > 1{
				total += connected[0].Value*connected[1].Value
			}
		}
	}
	return total
}

func processLine(line string)(*Line, error){
	foundSymbols := []*Symbol{}
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
			foundSymbols = append(foundSymbols, &Symbol{Symbol: charStr, Index: i})
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
			foundSymbols = append(foundSymbols, &Symbol{Symbol: charStr, Index: i})
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
