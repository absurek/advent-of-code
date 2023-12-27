package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	ID     int
	rounds []Round
}

func NewGame(line string) Game {
	parts := strings.Split(line, ":")
	strRounds := strings.Split(parts[1], ";")

	strid := strings.Split(strings.Trim(parts[0], " "), " ")[1]
	id, _ := strconv.Atoi(strid)

	var rounds []Round
	for _, round := range strRounds {
		var current Round
		cubes := strings.Split(round, ",")
		for _, cube := range cubes {
			data := strings.Split(strings.Trim(cube, " "), " ")
			switch data[1] {
			case "red":
				current.Red, _ = strconv.Atoi(data[0])
			case "green":
				current.Green, _ = strconv.Atoi(data[0])
			case "blue":
				current.Blue, _ = strconv.Atoi(data[0])
			}
		}

		rounds = append(rounds, current)
	}

	return Game{
		ID:     id,
		rounds: rounds,
	}
}

func NewGames(path string) []Game {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var games []Game
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		games = append(games, NewGame(scanner.Text()))
	}

	return games
}

func bounds(game Game, red int, green int, blue int) int {
	for _, round := range game.rounds {
		if round.Red > red {
			return 0
		}

		if round.Green > green {
			return 0
		}

		if round.Blue > blue {
			return 0
		}
	}

	return game.ID
}

func main() {
	games := NewGames("./assets/02.txt")
	sum := 0
	for _, game := range games {
		sum += bounds(game, 12, 13, 14)
	}

	fmt.Println("sum: ", sum)
}
