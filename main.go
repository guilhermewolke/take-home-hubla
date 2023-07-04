package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/guilhermewolke/take-home/cmd/database"
	"github.com/guilhermewolke/take-home/handlers"
)

func main() {
	http.HandleFunc("/setup-database", setup)
	http.HandleFunc("/upload", handlers.Upload)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func setup(w http.ResponseWriter, r *http.Request) {
	if err := database.Setup(false); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Erro ao rodar o setup: %s", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tabelas criadas com sucesso!"))
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("index.html").ParseFiles("web/index.html"))
	err := t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Erro ao carregar a página inicial: %s", err)))
		return
	}
}
