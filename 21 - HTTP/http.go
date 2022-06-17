// HTTP = Protocolo de comunicacao WEB. cliente - servidor

// Request - Response

// Rotas
// URI - identificador de recursos
// Métodos - GET, POST, PUT, DELETE
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Testing..."))
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregando os usuarios"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
