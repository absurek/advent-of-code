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

func power(game Game) int {
	red := 0
	green := 0
	blue := 0
	for _, round := range game.rounds {
		if round.Red > red {
			red = round.Red
		}

		if round.Green > green {
			green = round.Green
		}

		if round.Blue > blue {
			blue = round.Blue
		}
	}

	return red * green * blue
}

func main() {
	games := NewGames("./assets/02.txt")
	sum := 0
	for _, game := range games {
		sum += power(game)
	}

	fmt.Println("sum: ", sum)
}
