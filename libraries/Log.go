package libraries

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func WriteLogs(message map[string]interface{}) error {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening log file:", err)
	}
	defer logFile.Close()

	if len(message) == 0 {
		log.Println("No message to log")
		return err
	}

	log.SetOutput(io.MultiWriter(os.Stderr, logFile))

	var logMessage string
	for key, value := range message {
		logMessage += fmt.Sprintf("%s:%v ", key, value)
	}
	log.Println(strings.TrimSpace(logMessage))

	return err
}
