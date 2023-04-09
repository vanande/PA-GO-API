package db

import "fmt"

func Main() {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(50) NOT NULL, email VARCHAR(50) NOT NULL, age INT NOT NULL, city VARCHAR(50) NOT NULL)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO users(name, email, age, city) VALUES (?, ?, ?, ?)", "John Doe", "john@example.com", 30, "New York")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO users(name, email, age, city) VALUES (?, ?, ?, ?)", "Jane Doe", "jane@example.com", 25, "San Francisco")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO users(name, email, age, city) VALUES (?, ?, ?, ?)", "Bob Smith", "bob@example.com", 40, "London")
	if err != nil {
		panic(err)
	}

	fmt.Println("Table created and rows inserted successfully.")
}
