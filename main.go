package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"

// ...

func faq(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var db *sql.DB

		// Get the question from the C program
		body, err := ioutil.ReadAll(req.Body)
		question := string(body)
		fmt.Println("Question received: ", question)

		// Exit if the question is empty
		if question == "" {
			return
		}

		// Si la question contient des mots grossiers, on répond "Pas de gros mots"
		if question == "merde" {
			fmt.Fprintf(w, "Pas de gros mots")
			return
		}

		// Connect to the database
		db = connect(db)

		// Query the database
		rows, err := db.Query("SELECT answer FROM faq WHERE question = ?", question)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var answer string

		// Check if any rows were returned from the first query
		if !rows.Next() {
			// If no rows were returned, perform the second query to get all the questions in the database
			rows, err = db.Query("SELECT question FROM faq")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			var questions string
			for rows.Next() {
				err := rows.Scan(&questions)
				if err != nil {
					log.Fatal(err)
				}
				answer += questions + "\n"
			}
			if answer == "" {
				answer = "No match found, try this\n"
			}
		} else {
			// Get the answer from the database
			err := rows.Scan(&answer)
			if err != nil {
				log.Fatal(err)
			}
			// Send the answer to the C program
			fmt.Fprintf(w, answer)
		}
		fmt.Println("Answer sent: ", answer)
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

	db.Exec(`INSERT INTO faq(question, answer) VALUES ("How does it work?", "You pay me and I do the job")`)
	db.Exec(`INSERT INTO faq(question, answer) VALUES ("Is it a scam?", "Realistically, no.")`)
	db.Exec(`INSERT INTO faq(question, answer) VALUES ("Who are you?", "We are a team of 3 people.")`)
	db.Exec(`INSERT INTO faq(question, answer) VALUES ("Is Marc in your team?", "Yes and she is the best!")`)

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
