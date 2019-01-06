package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	//numero aleatorio a adivinar
	target := rand.Intn(100) + 1

	// fmt.Println(target)
	fmt.Println("he elegido un número entre 1 y 100")
	fmt.Println("Podes adivinarlo?")

	reader := bufio.NewReader(os.Stdin)

	// bucle
	acertar := false
	for i := 0; i < 10; i++ {
		fmt.Println("Te quedan", 10-i, "respuestas")

		fmt.Println("Adivina el número")

		entrada, error := reader.ReadString('\n')

		if error != nil {
			log.Fatal(error)
		}

		entrada = strings.TrimSpace(entrada)
		// numero ingreado
		adivina, error := strconv.Atoi(entrada)

		if error != nil {
			log.Fatal(error)
		}

		// fmt.Println(adivina)

		if adivina < target {
			fmt.Println("Es muy bajo")
		} else if adivina > target {
			fmt.Println("Es muy alto")
		} else {
			acertar = true
			fmt.Printf("Acertaste")
			break
		}
	}

	if !acertar {
		fmt.Print("No acertaste, el numero era", target)
	}
}
