package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

const redCubes = 12
const blueCubes = 14
const greenCubes = 13

type Game struct {
	gameID       int
	redCubeSet   []int
	blueCubeSet  []int
	greenCubeSet []int
}

func (g *Game) isGamePossible() bool {
	return !hasGreaterElement(g.redCubeSet, redCubes) &&
		!hasGreaterElement(g.blueCubeSet, blueCubes) &&
		!hasGreaterElement(g.greenCubeSet, greenCubes)
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	totalSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game, err := processGame(line)
		if err != nil {
			fmt.Println(err)
			return
		}

		if game.isGamePossible() {
			totalSum += game.gameID
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(totalSum)
}

func processGame(line string) (*Game, error) {
	game := Game{}
	splitText := strings.Split(line, ": ")

	gameIDStr := strings.Split(splitText[0], " ")[1]
	gameIDInt, err := strconv.Atoi(gameIDStr)
	if err != nil {
		return nil, err
	}
	game.gameID = gameIDInt

	gameSets := strings.Split(splitText[1], ";")
	for _, set := range gameSets {
		for _, cube := range strings.Split(set, ", ") {
			parts := strings.Split(strings.TrimSpace(cube), " ")
			cubeAmountStr, cubeColor := parts[0], parts[1]
			cubeAmountInt, err := strconv.Atoi(cubeAmountStr)
			if err != nil {
				return nil, err
			}

			switch cubeColor {
			case "red":
				game.redCubeSet = append(game.redCubeSet, cubeAmountInt)
			case "blue":
				game.blueCubeSet = append(game.blueCubeSet, cubeAmountInt)
			case "green":
				game.greenCubeSet = append(game.greenCubeSet, cubeAmountInt)
			}
		}
	}

	return &game, nil
}

func hasGreaterElement(arr []int, value int) bool {
	for _, elem := range arr {
		if elem > value {
			return true
		}
	}
	return false
}
