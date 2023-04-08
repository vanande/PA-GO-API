package db

import (
	"database/sql"
	"fmt"
)

// InitDB initializes the database connection.
func InitDB() (*sql.DB, error) {
	config, err := LoadConfig("C:\\Users\\me\\GolandProjects\\TogetherAndStronger\\routes\\db\\init\\config.json")
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
