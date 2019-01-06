package princiapl

import (
	"fmt"
	"log"
	"net/http"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

func iniciarServidor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor creado")
}

func main() {
	http.HandleFunc("/", iniciarServidor)
	error := http.ListenAndServe(connHost+":"+connPort, nil)

	if error != nil {
		log.Fatal("error al iniciar el server: ", error)
		return
	}
}
