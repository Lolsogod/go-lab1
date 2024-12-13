package file

import (
	"fmt"
	"os"
)

func CreateNumbersFile(n int, fileName string) {
	file, _ := os.Create(fileName)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file", err)
		}
	}(file)

	for i := 1; i <= n; i++ {
		_, err := file.WriteString(fmt.Sprintf("%d\n", i))
		if err != nil {
			fmt.Println("Error writing to file: ", err)
		}
	}

	fmt.Printf("Файл '%s' успешно создан с числами от 1 до %d\n", fileName, n)
}
