package handlers

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	"github.com/guilhermewolke/take-home/config"
	"github.com/guilhermewolke/take-home/internal/upload"
	"github.com/guilhermewolke/take-home/utils"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	log.Printf("handlers.Upload - Início do método")

	if r.Method != http.MethodPost {
		utils.SetFlashSession(w, r, false, "Método inválido")
		log.Printf("handlers.Upload - Método inválido")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	db, err := config.DBConnect()

	if err != nil {
		log.Printf("handlers.Upload - Erro ao se conectar com o banco: %s", err)
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("Error on connecting to database"))
		return
	}

	defer db.Close()
	//Receiving the file...
	log.Printf("handlers.Upload - lendo o arquivo")
	file, _, err := r.FormFile("file")

	if err != nil {
		log.Printf("handlers.Upload - Erro ao fazer o upload do arquivo: %s", err)
		utils.SetFlashSession(w, r, false, fmt.Sprintf("Erro ao fazer o upload do arquivo: %s", err))
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	defer file.Close()

	//Reading the file data...
	scanner := bufio.NewScanner(file)
	if err != nil {
		log.Printf("handlers.Upload - Erro ao fazer o upload do arquivo: %s", err)
		log.Printf("handlers.Upload - Erro ao ler o arquivo: %s", err)
		utils.SetFlashSession(w, r, false, fmt.Sprintf("Erro ao ler o arquivo: %s", err))
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	//Processing line by line...
	var lineCount int = 1
	for scanner.Scan() {
		log.Printf("handlers.Upload - processando linha do arquivo: '%s'", scanner.Text())
		if err = upload.ProcessLine(db, scanner.Text(), lineCount); err != nil {
			log.Printf("handlers.Upload - Erro ao ler as linhas do arquivo: %s", err)
			utils.SetFlashSession(w, r, false, fmt.Sprintf("Erro ao ler as linhas do arquivo: %s", err))
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		log.Printf("handlers.Upload - linha %d processada com sucesso!", lineCount)
		lineCount++
	}

	log.Printf("handlers.Upload - Fim do método")
	utils.SetFlashSession(w, r, true, "Importação concluída com sucesso!")
	http.Redirect(w, r, "/", http.StatusFound)
}
