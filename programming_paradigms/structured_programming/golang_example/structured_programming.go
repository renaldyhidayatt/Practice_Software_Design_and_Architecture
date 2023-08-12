package golang_example

import "fmt"

type structure interface {
	tambah(a int, b int) int
	kurang(a int, b int) int
	kali(a int, b int) int
}

type structure_class struct {
}

func (c structure_class) tambah(a int, b int) int {
	return a + b
}

func (c structure_class) kurang(a int, b int) int {
	return a - b
}

func (c structure_class) kali(a int, b int) int {
	return a * b
}

func main() {
	var calculator structure

	calculator = structure_class{}

	angka1 := 10
	angka2 := 5

	fmt.Println("Hasil Penjumlahan:", calculator.tambah(angka1, angka2))
	fmt.Println("Hasil Pengurangan:", calculator.kurang(angka1, angka2))
	fmt.Println("Hasil Perkalian:", calculator.kali(angka1, angka2))
}
