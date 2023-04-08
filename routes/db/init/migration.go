package db

import "fmt"

func Main() {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(50) NOT NULL, email VARCHAR(50) NOT NULL)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO users(name, email) VALUES (?, ?)", "John Doe", "john@example.com")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO users(name, email) VALUES (?, ?)", "Jane Doe", "jane@example.com")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO users(name, email) VALUES (?, ?)", "Bob Smith", "bob@example.com")
	if err != nil {
		panic(err)
	}

	fmt.Println("Table created and rows inserted successfully.")
}
