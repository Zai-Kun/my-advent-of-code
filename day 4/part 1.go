package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

const inputFile = "input.txt"

type Card struct{
	WinningNumbers []int
	MyNumbers      []int
}

func (c *Card) getCardPoints()int{
	totalPoints := 0
	for _, myNum := range c.MyNumbers{
		if elemInArray(c.WinningNumbers, myNum){
			if totalPoints > 0 {
				totalPoints *= 2
				continue
			}

			totalPoints = 1
		}
	}

	return totalPoints
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	totalPoints := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		card := processCard(line)
		totalPoints += card.getCardPoints()
	}

	fmt.Println(totalPoints)
}

func processCard(line string) *Card{
	allNumbers := strings.Split(line, ":")[1]
	winningNumbers, myNumbers := splitNumbers(allNumbers)

	return &Card{WinningNumbers: winningNumbers, MyNumbers: myNumbers}
}

func splitNumbers(allNumbers string)([]int, []int){
	tmpSplit := strings.Split(allNumbers, "|")
	winningNumbers := convertArrayToInt(strings.Fields(tmpSplit[0]))
	myNumbers := convertArrayToInt(strings.Fields(tmpSplit[1]))

	return winningNumbers, myNumbers
}

func convertArrayToInt(arr []string) []int{
	result := []int{}
	for _, elem := range arr{
		result = append(result, toInt(elem))
	}

	return result
}

func toInt(numStr string) int {
	convertedNum, _ := strconv.Atoi(numStr)
	return convertedNum
}

func elemInArray(arr []int, elem int) bool {
	for _, arrElem := range arr{
		if arrElem == elem{
			return true
		}
	}

	return false
}