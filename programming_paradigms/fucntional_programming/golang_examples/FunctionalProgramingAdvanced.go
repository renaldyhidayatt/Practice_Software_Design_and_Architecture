package golangexamples

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 22},
		{"David", 28},
		{"Eve", 21},
	}

	var namesOfAdults []string
	for _, person := range people {
		if person.Age >= 18 {
			namesOfAdults = append(namesOfAdults, person.Name)
		}
	}

	fmt.Println("People:", people)
	fmt.Println("Names of adults:", namesOfAdults)
}
