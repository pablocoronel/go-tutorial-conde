package main

import "fmt"

func main() {
	var miInt int
	var miPunteroInt *int

	miPunteroInt = &miInt

	fmt.Println(miPunteroInt)

	//declaracion corta
	var miBool bool
	miPunteroBool := &miBool

	fmt.Println(miPunteroBool)
}
