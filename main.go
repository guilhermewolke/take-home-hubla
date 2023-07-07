package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/guilhermewolke/take-home/config"
	"github.com/guilhermewolke/take-home/handlers"
	transactionDB "github.com/guilhermewolke/take-home/repository/transaction"
	"github.com/guilhermewolke/take-home/utils"
)

func main() {
	http.HandleFunc("/upload", handlers.Upload)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db, err := config.DBConnect()

	if err != nil {
		log.Printf("index - Erro ao se conectar com o banco: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error on connecting to database"))
		return
	}

	t := template.Must(template.New("index.html").ParseFiles("web/index.html"))
	context := make(map[string]interface{}, 0)

	tdb := transactionDB.NewTransactionDB(db)

	context["message"], context["status"] = utils.GetFlashSession(w, r)

	transactions, err := tdb.ListTransactions()

	if err != nil {
		context["errorMsg"] = "Ocorreu um erro ao listar as transações"
		log.Printf("index - Ocorreu um erro ao listar as transações: %s", err)
	}

	context["transactions"] = transactions

	err = t.Execute(w, context)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Erro ao carregar a página inicial: %s", err)))
		return
	}
}
