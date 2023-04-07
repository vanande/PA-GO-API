package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	_ "strings"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

// ...

func printProgress(ch chan int) {
	for i := 0; i < 100; i += 10 {
		ch <- i
		time.Sleep(time.Second)
	}
	ch <- 100
}
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

func faq(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		//var db *sql.DB

		// Read the question from the request body
		body, _ := ioutil.ReadAll(req.Body) // @TODO find less barbaric
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

		// Si la question contient des mots grossiers, on répond "Pas de gros mots"

		badword := "(?i)merde" // the (?i) makes the regex case-insensitive

		re := regexp.MustCompile(badword)

		match := re.MatchString(question)

		if match {
			fmt.Println("Bad words detected!")
			fmt.Fprintf(w, "Bad words detected!")
			// @TODO add it
		}

		// Send the question to the GRC server

		response, err := askToGrcServer(question)
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

func main() {
	// Start the server
	http.HandleFunc("/", faq)
	http.ListenAndServe(":9000", nil)
}

func askToGrcServer(question string) (string, error) {
	//@TODO turn grc into endpoint
	viper.SetConfigType("properties")
	viper.SetConfigFile("config.conf")
	err := viper.ReadInConfig()
	if err != nil {
		return "", fmt.Errorf("Failed to read conf file: %s", err)
	}

	username := viper.GetString("username")
	password := viper.GetString("password")

	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", "192.168.237.129:22", sshConfig)
	if err != nil {
		return "", fmt.Errorf("Failed to dial: %s", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return "", fmt.Errorf("Failed to create session: %s", err)
	}
	defer session.Close()

	cmd := fmt.Sprintf("cd go/verbaflow/ && echo '\nQ: %s \n\nA:' | /usr/local/go/bin/go run ./examples/prompttester --dconfig ./examples/prompttester/config.yaml", question)

	out, err := session.Output(cmd)
	if err != nil {
		return "", fmt.Errorf("Failed to run command: %s", err)
	}

	return string(out), nil
}

func askFlaskServer(question string) (string, error) {
	client := &http.Client{}

	reqBody := []byte(fmt.Sprintf(`{"question": "%s"}`, question))
	req, err := http.NewRequest("POST", "http://127.0.0.1:5001/faq", bytes.NewBuffer(reqBody))
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
	json.Unmarshal(body, &jsonResponse)

	return jsonResponse["answer"], nil
}
