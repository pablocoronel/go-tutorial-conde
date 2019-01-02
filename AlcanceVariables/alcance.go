package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Pide el ingreso
	fmt.Print("Ingresa una puntuacion: ")

	// almacena el ingreso por teclado
	reader := bufio.NewReader(os.Stdin)

	//convierte a string lo almacenado
	entrada, error := reader.ReadString('\n')

	// Si hay error, lo muestra (nil es el valor cuando No hay error)
	if error != nil {
		log.Fatal(error)
	}

	//Quita espacios de la entrada
	entrada = strings.TrimSpace(entrada)

	//Convierte el testo ingresado a Float64 solo si es posible: "10" si, "Pablo" no
	puntuacion, error := strconv.ParseFloat(entrada, 64)

	// Si hay error, lo muestra (nil es el valor cuando No hay error)
	if error != nil {
		log.Fatal(error)
	}

	//Declara la variable de respuesta
	var status string

	//Condicional para evaluar el valor ingresado
	if puntuacion >= 50 {
		status = "Aprobado"
	} else {
		status = "Desaprobado"
	}

	//Muestra el texto final
	fmt.Println("Una puntuacion de:", puntuacion, "es", status)

}
