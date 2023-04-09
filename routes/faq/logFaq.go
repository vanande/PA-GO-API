package faq

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func logQuestionAnswer(question, answer string) error {
	cwd, _ := os.Getwd()

	logFilePath := filepath.Join(cwd, "faq.log")
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		file, _ := os.Create(logFilePath)
		file.Close()
	}

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open log file: %s", err)
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("%s | Question: %s | Answer: %s\n", timestamp, question, answer)

	if _, err := file.WriteString(logMessage); err != nil {
		return fmt.Errorf("Failed to write log message to file: %s", err)
	}

	return nil
}
