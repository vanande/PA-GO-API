package faq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
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

func Faq(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		body, _ := io.ReadAll(req.Body) // @TODO find less barbaric
		question := string(body)
		fmt.Println("Question received: ", question)

		if question == "" {
			return
		}

		if len(question) > 70 {
			fmt.Println("Question too long")
			fmt.Fprintf(w, "Question too long")
			return
		}

		badword := "(?i)merde" // the (?i) makes the regex case-insensitive

		re := regexp.MustCompile(badword)

		match := re.MatchString(question)

		if match {
			fmt.Println("Bad words detected!")
			fmt.Fprintf(w, "Bad words detected!")
			// @TODO add it to the log
		}

		response, err := askFlaskServer(question)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Fprintf(w, "Error: %s", err)
		}
		fmt.Println("Response: ", response)
		fmt.Fprintln(w, "Response: ", response)

		err = logQuestionAnswer(question, response)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Fprintf(w, "Error: %s", err)
		}
	}
}

func askFlaskServer(question string) (string, error) {
	client := &http.Client{}

	reqBody := []byte(fmt.Sprintf(`{"question": "%s"}`, question))
	req, err := http.NewRequest("POST", "http://127.0.0.1:5000/api", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("Failed to create request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to send request: %s", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var jsonResponse map[string]string
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return "", fmt.Errorf("Failed to parse JSON response: %s", err)
	}

	return jsonResponse["answer"], nil
}
