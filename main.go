package main

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net/http"
	"regexp"
)
import _ "github.com/go-sql-driver/mysql"

// ...

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

	}
}

func main() {
	// Start the server
	http.HandleFunc("/", faq)
	http.ListenAndServe(":9000", nil)
}
func askToGrcServer(question string) (string, error) {
	// Load the conf file
	viper.SetConfigType("properties")
	viper.SetConfigFile("config.conf")
	err := viper.ReadInConfig()
	if err != nil {
		return "", fmt.Errorf("Failed to read conf file: %s", err)
	}

	// Get values from the conf file
	username := viper.GetString("username")
	password := viper.GetString("password")

	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", "192.168.244.131:22", sshConfig)
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
