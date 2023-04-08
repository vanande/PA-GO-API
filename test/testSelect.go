package main

import (
	"fmt"
	"log"

	"TogetherAndStronger/routes/db"
)

func main() {
	tableName := "users"
	columns := []string{"id", "name", "email"}
	conditions := map[string]interface{}{}

	rows, err := db.SelectQuery(tableName, columns, conditions)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email string
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d, name: %s, email: %s\n", id, name, email)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
