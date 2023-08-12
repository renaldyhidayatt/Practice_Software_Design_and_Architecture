package golangexamples

import "fmt"

type Student struct {
	Name      string
	Age       int
	StudentID string
}

func (s Student) DisplayInfo() {
	fmt.Printf("Name: %s, Age: %d, Student ID: %s\n", s.Name, s.Age, s.StudentID)
}

func main() {
	student := Student{Name: "Alice", Age: 20, StudentID: "12345"}
	student.DisplayInfo()
}
