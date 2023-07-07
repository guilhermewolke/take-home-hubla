package utils

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("take-home"))

// Criação de flash-sessions para mensagens de erro e sucesso
func SetFlashSession(w http.ResponseWriter, r *http.Request, success bool, message string) {
	session, err := Store.Get(r, "flash-session")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.AddFlash(message, "message")

	if success {
		session.AddFlash("success", "status")
	} else {
		session.AddFlash("failure", "status")
	}
	session.Save(r, w)
}

// Recuperação de flash-sessions para impressão de mensagens de erro e sucesso na tela de listagem de lançamentos
func GetFlashSession(w http.ResponseWriter, r *http.Request) (string, string) {
	session, err := Store.Get(r, "flash-session")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", ""
	}

	message := session.Flashes("message")
	status := session.Flashes("status")

	m := fmt.Sprintf("%v", message)
	s := fmt.Sprintf("%v", status)

	session.Save(r, w)
	return m, s
}
