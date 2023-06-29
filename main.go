package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/guilhermewolke/take-home/cmd/database"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/setup-database", setup)

	http.ListenAndServe(":8080", mux)
}

func setup(w http.ResponseWriter, r *http.Request) {
	c := context.Background()
	if err := database.Setup(c); err != nil {
		panic(err)
	}
	w.WriteHeader(200)
	w.Write([]byte("Tabelas criadas com sucesso!"))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow ord")
}
