package libraries

import (
	"io"
	"log"
	"os"
)

func WriteLogs(message map[string]interface{}) error {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening log file:", err)
	}
	defer logFile.Close()

	log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	log.Printf("Message: %s\n", message)
	return err
}
