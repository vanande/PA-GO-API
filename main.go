package main

import (
	"TogetherAndStronger/routes/auth"
	"TogetherAndStronger/routes/db/db_handler"
	"TogetherAndStronger/routes/faq"
	"TogetherAndStronger/routes/signup"
	"net/http"
	_ "strings"
)

func main() {
	//db.Main()

	// Start the server
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
	// go run ..\..\go\go1.20.1\src\crypto\tls\generate_cert.go -host="127.0.0.1"
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
	//http.ListenAndServeTLS(":9000", "cert.pem", "key.pem", nil)
}
