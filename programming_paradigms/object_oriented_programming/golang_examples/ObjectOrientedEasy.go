package golangexamples

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}

func (c Car) DisplayInfo() {
	fmt.Printf("Brand: %s, Model: %s, Year: %d\n", c.Brand, c.Model, c.Year)
}

func main() {
	car := Car{Brand: "Toyota", Model: "Camry", Year: 2022}
	car.DisplayInfo()
}
