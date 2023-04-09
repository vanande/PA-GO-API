package main

import (
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"log"
)

func main() {
	tableName := "users"
	columns := []string{"id", "name", "email", "age", "city"}
	conditions := map[string]interface{}{}

	rows, err := query.SelectQuery(tableName, columns, conditions)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email string
		var age int
		var city string
		err = rows.Scan(&id, &name, &email, &age, &city)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d, name: %s, email: %s, age: %d, city: %s \n", id, name, email, age, city)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
