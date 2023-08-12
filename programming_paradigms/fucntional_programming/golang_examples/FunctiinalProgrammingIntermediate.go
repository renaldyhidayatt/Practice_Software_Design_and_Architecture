package golangexamples

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}

	var filteredNames []string
	for _, name := range names {
		if len(name) > 4 {
			filteredNames = append(filteredNames, strings.ToUpper(name))
		}
	}

	fmt.Println("Original names:", names)
	fmt.Println("Filtered and uppercase names:", filteredNames)
}
