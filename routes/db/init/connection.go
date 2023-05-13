package db

import (
	"database/sql"
	"fmt"
)

// InitDB initializes the database connection.
func InitDB() (*sql.DB, error) {
<<<<<<< HEAD
	config, err := LoadConfig("C:\\Users\\me\\GolandProjects\\TogetherAndStronger\\routes\\db\\init\\config.json")
=======
	config, err := LoadConfig("C:\\Users\\me\\GolandProjects\\TogetherAndStronger\\routes\\db\\init\\config.json                        ")
>>>>>>> 1784635def712a3f39b341ecc4c2c2f38ab4d056
	if err != nil {
		return nil, err
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		config.Username, config.Password, config.Hostname, config.Port)

	db, err := sql.Open(config.Dialect, connectionString)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.Name))
	if err != nil {
		return nil, err
	}

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Username, config.Password, config.Hostname, config.Port, config.Name)

	db, err = sql.Open(config.Dialect, connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to database.")
	return db, nil
}
