package main

import (
	"awesomeProject/internal/benchmark"
	"awesomeProject/internal/calculator"
	"awesomeProject/internal/config"
	"awesomeProject/internal/file"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации:", err)
	}

	if cfg.MakeNewFile {
		file.CreateNumbersFile(cfg.NumberAmount, cfg.FileName)
	}

	numbers := file.ReadNumbersFromFile(cfg.FileName)
	fmt.Printf("Всего чисел в файле: %d\n", len(numbers))

	if cfg.BenchMode {
		benchmark.Run(cfg.OperationName, cfg.Multiplier, numbers)
	} else {
		results, _ := calculator.Process(
			cfg.OperationName,
			cfg.Multiplier,
			cfg.GoRoutineAmount,
			cfg.ThreadsAmount,
			numbers,
		)

		if cfg.WriteToFile {
			file.WriteResultsToFile(results, cfg.FileName)
		}
	}

	fmt.Println("многопоточная обработка успешно завершена")
}
