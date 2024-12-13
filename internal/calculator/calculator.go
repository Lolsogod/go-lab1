package calculator

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Process(operationName string, multiplier, threadsAmount int, goRoutinesAmount int, numbers []int) ([][]int, time.Duration) {
    runtime.GOMAXPROCS(threadsAmount)
    
    chunks := make([][]int, goRoutinesAmount)
    results := make([][]int, goRoutinesAmount)
    chunkSize := len(numbers) / goRoutinesAmount
    
    for i := 0; i < goRoutinesAmount; i++ {
        start := i * chunkSize
        end := (i + 1) * chunkSize
        if i == goRoutinesAmount-1 {
            end = len(numbers)
        }
        chunks[i] = numbers[start:end]
        results[i] = make([]int, len(chunks[i]))
    }
    
    var wg sync.WaitGroup
    timeNow := time.Now()
    
    for i := 0; i < goRoutinesAmount; i++ {
        wg.Add(1)
        go func(threadNum int) {
            defer wg.Done()
            for j, num := range chunks[threadNum] {
                results[threadNum][j] = calculate(operationName, num, multiplier)
            }

            fmt.Printf("Горутина %d завершила работу, обработано чисел: %d\n", threadNum, len(chunks[threadNum]))
        }(i)
    }
    
    wg.Wait()
	
	fmt.Println("----------------------------------------------------")
	fmt.Printf("Размер чанка: %d\n", chunkSize)
	fmt.Printf("Количество потоков: %d\n", threadsAmount)
	fmt.Println("----------------------------------------------------")
	fmt.Printf("Для %d потоков прошло: %d\n", threadsAmount, time.Since(timeNow))
	fmt.Println("----------------------------------------------------")

	return results, time.Since(timeNow)
}
