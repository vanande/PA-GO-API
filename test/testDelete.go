package main

import (
	"TogetherAndStronger/routes/db"
	_ "strings"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

	db.DeleteQuery("users", "id = ?", 1)
}
