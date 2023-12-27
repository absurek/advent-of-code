package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func decode(line string) int {
	var digits []int
	for _, char := range line {
		digit, err := strconv.Atoi(string(char))
		if err == nil {
			digits = append(digits, digit)
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
