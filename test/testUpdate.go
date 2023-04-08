package main

import (
	"TogetherAndStronger/routes/db"
)

func main() {
	// Define table name and update data
	tableName := "users"
	updateData := map[string]interface{}{
		"name": "Jane",
	}

	// Define where clause conditions
	conditions := map[string]interface{}{
		"id": 5,
	}

	// Call update function
	err := db.UpdateQuery(tableName, updateData, conditions)
	if err != nil {
		panic(err)
	}
}
