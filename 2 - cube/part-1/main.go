package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	total := 0

	games := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	for i, game := range games {

		maxCubesPerRound := map[string]int{
			"red":   12,
			"blue":  14,
			"green": 13,
		}

		gameId := i + 1
		gameResultStartingIndex := strings.Index(game, ": ") + 2

		cubesInGame := game[gameResultStartingIndex:]

		cubeSets := strings.Split(cubesInGame, ";")

		isValidRound := true

		for _, set := range cubeSets {
			cubeResults := strings.Split(set, ",")
			for _, result := range cubeResults {
				trimmedResult := strings.TrimSpace(result)
				amount, _ := strconv.Atoi(strings.Split(trimmedResult, " ")[0])
				color := strings.Split(trimmedResult, " ")[1]

				if amount > maxCubesPerRound[color] {
					isValidRound = false
				}
			}
		}

		if isValidRound {
			total += gameId
		}

	}

	fmt.Println(total)

}
