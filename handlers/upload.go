package handlers

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	"github.com/guilhermewolke/take-home/config"
	"github.com/guilhermewolke/take-home/internal/upload"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	log.Printf("handlers.Upload - Início do método")

	if r.Method != http.MethodPost {
		log.Printf("handlers.Upload - Método inválido")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Invalid method"))
		return
	}

	db, err := config.DBConnect()

	if err != nil {
		log.Printf("handlers.Upload - Erro ao se conectar com o banco: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error on connecting to database")))
		return
	}

	defer db.Close()
	//Receiving the file...
	log.Printf("handlers.Upload - lendo o arquivo")
	file, _, err := r.FormFile("file")

	if err != nil {
		log.Printf("handlers.Upload - Erro ao fazer o upload do arquivo: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error on uploading file")))
		return
	}

	defer file.Close()

	//Reading the file data...
	scanner := bufio.NewScanner(file)
	if err != nil {
		log.Printf("handlers.Upload - Erro ao ler o arquivo: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error on reading uploaded file")))
		return
	}

	//Processing line by line...
	for scanner.Scan() {
		log.Printf("handlers.Upload - processando linha do arquivo: '%s'", scanner.Text())
		if err = upload.ProcessLine(db, scanner.Text()); err != nil {
			log.Printf("handlers.Upload - Erro ao ler as linhas do arquivo: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error on reading the file lines")))
			return
		}
		log.Printf("handlers.Upload - linha processada com sucesso!")
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Upload finished with success!"))
	log.Printf("handlers.Upload - Fim do método")
}
