package handlers

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/guilhermewolke/take-home/internal/upload"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Invalid method"))
	}
	//Receiving the file...
	file, _, err := r.FormFile("file")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error on uploading file: %s", err)))
	}

	defer file.Close()

	//Reading the file data...
	scanner := bufio.NewScanner(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error on reading uploaded file: %s", err)))
	}

	for scanner.Scan() {
		if err = upload.ProcessLine(scanner.Text()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error on reading the file lines: %s", err)))
		}
	}

	//Processing line by line...
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Upload finished with success!"))
}
