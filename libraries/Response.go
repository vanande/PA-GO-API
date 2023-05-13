package libraries

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func Response(w http.ResponseWriter, message map[string]interface{}, code int) {

	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening log file:", err)
	}
	defer logFile.Close()

	msg, OK := message["message"].(string)
	if !OK {
		msg = "No message provided"
	}

	log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	log.Printf("Response - Code: %d - Message: %s\n", code, msg)

	response := map[string]interface{}{}
	for key, value := range message {
		response[key] = value
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8	")
	w.WriteHeader(code)
	w.Write(jsonResponse)
}
