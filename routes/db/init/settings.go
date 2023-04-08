package db

import (
	"encoding/json"
	"os"
)

type Config struct {
	Dialect  string
	Hostname string
	Port     int
	Name     string
	Username string
	Password string
}

// LoadConfig loads the configuration from a JSON file. Here is an example of JSON file :
//
//	{
//	 "Dialect": "mysql",
//	 "Hostname": "localhost",
//	 "Port": 3306,
//	 "Name": "TaS", -> name of database
//	 "Username": "root",
//	 "Password": "root"
//	}
func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
