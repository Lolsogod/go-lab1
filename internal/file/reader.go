package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNumbersFromFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			continue
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error parsing file:", err)
		return nil
	}

	return numbers
}
