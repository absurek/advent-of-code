package main

import (
	"fmt"
	"strconv"
	"strings"
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
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		sum += decode(line)
	}

	fmt.Println("sum: ", sum)
}
