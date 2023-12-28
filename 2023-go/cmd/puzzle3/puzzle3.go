package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Number struct {
	Value     int
	Positions []Position
}

type Position struct {
	X int
	Y int
}

func getFieldValue(field [][]bool, x int, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if y >= len(field) {
		return false
	}

	if x >= len(field[y]) {
		return false
	}

	return field[y][x]
}

func (p *Position) hasAdjacentSymbol(field [][]bool) bool {
	return getFieldValue(field, p.X+1, p.Y) || getFieldValue(field, p.X-1, p.Y) ||
		getFieldValue(field, p.X, p.Y+1) || getFieldValue(field, p.X, p.Y-1) ||
		getFieldValue(field, p.X+1, p.Y+1) || getFieldValue(field, p.X+1, p.Y-1) ||
		getFieldValue(field, p.X-1, p.Y+1) || getFieldValue(field, p.X-1, p.Y-1)
}

func isDigit(char rune) bool {
	switch char {
	case '0':
		return true
	case '1':
		return true
	case '2':
		return true
	case '3':
		return true
	case '4':
		return true
	case '5':
		return true
	case '6':
		return true
	case '7':
		return true
	case '8':
		return true
	case '9':
		return true
	default:
		return false
	}
}

func readLine(lineIndex int, line string) ([]Number, []bool) {
	row := []bool{}
	numbers := []Number{}
	digits := []rune{}
	positions := []Position{}
	for index, char := range line {
		if isDigit(char) {
			row = append(row, false)
			digits = append(digits, char)
			positions = append(positions, Position{X: index, Y: lineIndex})
		} else {
			value, err := strconv.Atoi(string(digits))
			if err == nil {
				numbers = append(numbers, Number{Value: value, Positions: positions})
			}
			digits = nil
			positions = nil

			if char == '.' {
				row = append(row, false)
			} else {
				row = append(row, true)
			}
		}
	}

	value, err := strconv.Atoi(string(digits))
	if err == nil {
		numbers = append(numbers, Number{Value: value, Positions: positions})
	}

	return numbers, row
}

func readFile(path string) ([]Number, [][]bool) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbers := []Number{}
	field := [][]bool{}
	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		nums, row := readLine(index, scanner.Text())
		field = append(field, row)
		numbers = append(numbers, nums...)
		index++
	}

	return numbers, field
}

func isPartNumber(number Number, field [][]bool) bool {
	for _, position := range number.Positions {
		if position.hasAdjacentSymbol(field) {
			return true
		}
	}

	return false
}

func partnum(number Number, field [][]bool) int {
	if isPartNumber(number, field) {
		return number.Value
	}

	return 0
}

func main() {
	numbers, field := readFile("./assets/03.txt")
	sum := 0
	for _, number := range numbers {
		sum += partnum(number, field)
	}

	fmt.Println("sum: ", sum)
}
