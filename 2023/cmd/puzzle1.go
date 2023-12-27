package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseDigit(digits []string, str string) int {
	for _, digit := range digits {
		idx := strings.Index(str, digit)
		if idx != -1 {
			return idx
		}
	}

	return -1
}

func findDigit(str string) (int, int) {
	index := parseDigit([]string{"one", "1"}, str)
	if index != -1 {
		return 1, index
	}

	index = parseDigit([]string{"two", "2"}, str)
	if index != -1 {
		return 2, index
	}

	index = parseDigit([]string{"three", "3"}, str)
	if index != -1 {
		return 3, index
	}

	index = parseDigit([]string{"four", "4"}, str)
	if index != -1 {
		return 4, index
	}

	index = parseDigit([]string{"five", "5"}, str)
	if index != -1 {
		return 5, index
	}

	index = parseDigit([]string{"six", "6"}, str)
	if index != -1 {
		return 6, index
	}

	index = parseDigit([]string{"seven", "7"}, str)
	if index != -1 {
		return 7, index
	}

	index = parseDigit([]string{"eight", "8"}, str)
	if index != -1 {
		return 8, index
	}

	index = parseDigit([]string{"nine", "9"}, str)
	if index != -1 {
		return 9, index
	}

	return 0, -1
}

func decode(line string) int {
	var digits []int
	var chars []rune
	for _, char := range line {
		chars = append(chars, char)
		digit, index := findDigit(string(chars))
		if index != -1 {
			digits = append(digits, digit)
			chars = chars[index+1:]
		}
	}

	return digits[0]*10 + digits[len(digits)-1]
}

func main() {
	file, err := os.Open("./assets/01-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += decode(scanner.Text())
	}

	fmt.Println("sum: ", sum)
}
