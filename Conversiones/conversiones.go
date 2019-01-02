package main

import "fmt"

func main() {
	precio := 100
	iva := 0.03
	tasa := float64(precio) * iva
	total := float64(precio) + iva

	fmt.Println("El precio es ", precio, "Euros")
	fmt.Println("La tasa es", tasa, "Euros")
	fmt.Println("El costo total es", total, "Euros")
}
