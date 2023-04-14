package faq

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

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

		fmt.Println("Question: ", question)
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
