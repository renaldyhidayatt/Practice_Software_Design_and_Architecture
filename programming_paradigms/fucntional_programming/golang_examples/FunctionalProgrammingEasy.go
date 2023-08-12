package golangexamples

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	squaredNumbers := make([]int, len(numbers))
	for i, n := range numbers {
		squaredNumbers[i] = n * n
	}

	fmt.Println("Original numbers:", numbers)
	fmt.Println("Squared numbers:", squaredNumbers)
}
