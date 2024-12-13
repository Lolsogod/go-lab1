package config

import (
	"os"
	"strconv"
)

type Config struct {
	NumberAmount    int
	Multiplier      int
	ThreadsAmount   int
	GoRoutineAmount int
	MakeNewFile     bool
	WriteToFile     bool
	BenchMode       bool
	FileName        string
	OperationName   string
}

func LoadConfig() (*Config, error) {
	config := &Config{
		NumberAmount:    getNumberFromEnv("NUMBER_AMOUNT"),
		Multiplier:      getNumberFromEnv("MULTIPLIER"),
		ThreadsAmount:   getNumberFromEnv("THREADS"),
		GoRoutineAmount: getNumberFromEnv("GOROUTINES"),
		MakeNewFile:     getBoolFromEnv("MAKE_NEW_FILE"),
		WriteToFile:     getBoolFromEnv("WRITE_TO_FILE"),
		BenchMode:       getBoolFromEnv("BENCH_MODE"),
	}

	config.FileName, _ = os.LookupEnv("FILE_NAME")
	config.OperationName, _ = os.LookupEnv("OPERATION_NAME")

	return config, nil
}

func getNumberFromEnv(key string) int {
	rawNumber, _ := os.LookupEnv(key)
	parsedNumber, _ := strconv.Atoi(rawNumber)
	return parsedNumber
}

func getBoolFromEnv(key string) bool {
	rawValue, _ := os.LookupEnv(key)
	parsedBool, _ := strconv.ParseBool(rawValue)
	return parsedBool
}
