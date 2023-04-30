package main

import (
	"TogetherAndStronger/routes/auth"
	"TogetherAndStronger/routes/db/db_handler"
	"TogetherAndStronger/routes/faq"
	"TogetherAndStronger/routes/lookfor"
	"TogetherAndStronger/routes/signup"
	"fmt"
	"net/http"
	_ "strings"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Best angular website of the century incoming... ")
	fmt.Println("Got one")
}

func main() {
	//db.Main()

	fmt.Println("Server starting...")

	// Start the server
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/faq", faq.Faq)
	http.HandleFunc("/db/create", db_handler.Create)
	http.HandleFunc("/db/select", db_handler.Select)
	http.HandleFunc("/db/update", db_handler.Update)
	http.HandleFunc("/db/delete", db_handler.Delete)
	http.HandleFunc("/auth/LoginCompany", auth.LoginCompany)
	http.HandleFunc("/auth/LoginPresta", auth.LoginPresta)
	http.HandleFunc("/auth/LoginSalary", auth.LoginSalary)
	http.HandleFunc("/signup/SignupSalary", signup.SignupSalary)
	http.HandleFunc("/signup/SignupCompany", signup.SignupCompany)
	http.HandleFunc("/signup/SignupPresta", signup.SignupPresta)
	http.HandleFunc("/lookfor/LookForSalary", lookfor.LookForSalary)
	// go run ..\..\go\go1.20.1\src\crypto\tls\generate_cert.go -host="127.0.0.1"
	http.ListenAndServe(":9000", nil)
	//err := http.ListenAndServeTLS(":9000", "cert.pem", "key.pem", nil)
}
