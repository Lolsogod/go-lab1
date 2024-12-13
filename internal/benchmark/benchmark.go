package benchmark

import (
	"awesomeProject/internal/calculator"
	"fmt"
	"time"
)

func Run(operationName string, multiplier int, numbers []int) {
	threadsToTest := []int{1, 2, 3, 4, 5, 6, 8, 10, 12, 20, 30, 100, 300, 500, 700, 1000, 2000, 5000, 10000}
	times := make(map[int]time.Duration)
	const iterations = 5 // количество прогонов для каждой конфигурации

	for _, threads := range threadsToTest {
		var totalTime time.Duration
		for i := 0; i < iterations; i++ {
			_, executionTime := calculator.Process(operationName, multiplier, threads, threads, numbers)
			totalTime += executionTime
		}
		times[threads] = totalTime / iterations
	}

	fmt.Println("\nBenchmark results (averaged over", iterations, "runs):")
	for _, threads := range threadsToTest {
		fmt.Printf("Threads - %d, Time - %d\n", threads, times[threads].Milliseconds())
	}
}
