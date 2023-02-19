package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)
import _ "github.com/go-sql-driver/mysql"

// ...

func faq(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var db *sql.DB

		// Get the question from the C program
		body, _ := ioutil.ReadAll(req.Body)
		question := string(body)
		fmt.Println("Question received: ", question)

		// Exit if the question is empty
		if question == "" {
			return
		}

		// Si la question contient des mots grossiers, on répond "Pas de gros mots"

		// If the question contains bad words, we answer "No bad words"
		// Define a regular expression pattern to match the word "merde"
		badword := "(?i)merde" // the (?i) makes the regex case-insensitive

		// Compile the regular expression
		re := regexp.MustCompile(badword)

		// Test whether a string matches the regular expression
		match := re.MatchString(question)

		if match {
			fmt.Println("Bad words detected!")
			fmt.Fprintf(w, "Bad words detected!")
		}
		// Connect to the database
		db = connect(db)

		// Define the keywords to look for
		keywords := []string{"how", "scam", "who", "Marc"}

		// Check if any of the keywords are in the input string
		var found bool = false
		for _, keyword := range keywords {
			if strings.Contains(strings.ToLower(question), strings.ToLower(keyword)) {
				// If a keyword is found, perform the first query to get the answer from the database
				found = true
				// Query the database
				query := fmt.Sprintf("SELECT answer FROM faq WHERE LOWER(question) LIKE '%%%s%%'", keyword)

				// Query the database
				rows, _ := db.Query(query)

				defer rows.Close()

				if rows.Next() {
					var answer string
					rows.Scan(&answer)
					fmt.Printf("Answer: %s\n", answer)
					fmt.Fprintf(w, "Answer: %s\n", answer)
				} else {
					fmt.Println("No matching answer found.")
					fmt.Fprintln(w, "No matching answer found.")
				}
				break
			}
		}
		fmt.Println("Found: ", found)

		if !found {
			fmt.Fprintf(w, "No matching answer found. Try these keywords: how, scam, who, Marc")
		}
	}
}

func connect(db *sql.DB) *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3308)/pa_example")
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}

func injectDataInDB(db *sql.DB) {
	// Connect to the database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3308)/pa_example")
	if err != nil {
		log.Fatal(err)
	}

	// Create database pa_example and table faq with the following columns : id, question, answer
	db.Exec(`CREATE TABLE IF NOT EXISTS faq (
		id INT NOT NULL AUTO_INCREMENT,
		question varchar(255) NOT NULL,
		answer text NOT NULL,
		PRIMARY KEY (id)
	) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=1`)

	if err != nil {
		panic(err.Error())
	}

	db.Exec(`INSERT INTO faq(question, answer) VALUES ("how does it work", "You pay me and I do the job")`)
	db.Exec(`INSERT INTO faq(question, answer) VALUES ("why are you doing this", "Because pay me and I do the job")`)
	db.Exec(`INSERT INTO faq(question, answer) VALUES ("is it a scam", "Realistically, no.")`)
	db.Exec(`INSERT INTO faq(question, answer) VALUES ("who are you", "We are a team of 3 people.")`)
	db.Exec(`INSERT INTO faq(question, answer) VALUES ("is Marc in your team", "Yes and she is the best!")`)

	fmt.Println("SQL statements executed successfully")
}

func main() {
	var db *sql.DB

	// Connect to the database and inject data
	db = connect(db)
	//injectDataInDB(db)

	// Start the server
	http.HandleFunc("/", faq)
	http.ListenAndServe(":9000", nil)
}
