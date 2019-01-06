package princiapl

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func iniciarServidor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor creado")
}

func main() {
	http.HandleFunc("/", iniciarServidor)
	error := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)

	if error != nil {
		log.Fatal("error al iniciar el server: ", error)
		return
	}
}
