package calculator

import (
	"fmt"
	"math"
)

var operations = map[string]func(x, y int) int{
    "multiply":  multiply,
    "power":     power,
    "factorial": func(x, y int) int { return factorial(x) },
    "fibonacci": func(x, y int) int { return fibonacci(x) },
}

func calculate(operationName string, x int, y int) int {
    operation, exists := operations[operationName]
    if !exists {
        fmt.Println("Unrecognized operation: " + operationName)
        return 0
    }
    return operation(x, y)
}

func multiply(x int, y int) int {
	return x * y
}

func power(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func factorial(n int) int {
    if n <= 0 {
        return 1
    }

    result := 1

    for i := 1; i <= n; i++ {
        result *= i
    }

    return result
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	prev, curr := 0, 1

	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}
