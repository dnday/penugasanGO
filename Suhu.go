package main

import (
	"fmt"
)

func main() {
	var celsius float64
	var choice int

	// Input suhu dalam Celsius
	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scanln(&celsius)

	// Pilihan konversi
	fmt.Println("Pilih satuan konversi:")
	fmt.Println("1. Fahrenheit")
	fmt.Println("2. Kelvin")
	fmt.Println("3. Reamur")
	fmt.Print("Masukkan pilihan (1/2/3): ")
	fmt.Scanln(&choice)

	// Switch-case untuk konversi
	switch choice {
	case 1:
		// Konversi ke Fahrenheit
		fahrenheit := (celsius * 9 / 5) + 32
		fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n", celsius, fahrenheit)
	case 2:
		// Konversi ke Kelvin
		kelvin := celsius + 273
		fmt.Printf("%.2f Celsius = %.2f Kelvin\n", celsius, kelvin)
	case 3:
		// Konversi ke Reamur
		reamur := celsius * 4 / 5
		fmt.Printf("%.2f Celsius = %.2f Reamur\n", celsius, reamur)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
