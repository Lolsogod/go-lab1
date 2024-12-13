package file

import (
	"bufio"
	"fmt"
	"os"
)

func WriteResultsToFile(results [][]int, filename string) {
	file, _ := os.Create(filename)
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, chunk := range results {
		for _, num := range chunk {
			_, err := writer.WriteString(fmt.Sprintf("%d\n", num))
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	err := writer.Flush()
	if err != nil {
		fmt.Println(err)
	}
}
