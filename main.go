package main

import (
	"TogetherAndStronger/routes/auth"
	"TogetherAndStronger/routes/db/db_handler"
	"TogetherAndStronger/routes/entity/salary"
	"TogetherAndStronger/routes/faq"
	"TogetherAndStronger/routes/list"
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

	fmt.Println("Server starting...")

	// Start the server
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./img"))))
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
	http.HandleFunc("/lookfor/LookForCompany", lookfor.LookForCompany)
	//	http.HandleFunc("/lookfor/LookForPresta", lookfor.LookForPresta)
	http.HandleFunc("/salary/getActivities", salary.GetSalaryActivities)
	http.HandleFunc("/salary/addInfo", salary.AddSalaryInfo)
	http.HandleFunc("/salary/deleteInfo", salary.DeleteSalaryInfo)
	http.HandleFunc("/salary/getInfos", salary.GetSalaryInfo)
	http.HandleFunc("/salary/addInfos", salary.AddSalaryInfo)
	http.HandleFunc("/list/infos", list.ListInfos)
	err := http.ListenAndServeTLS(":9000", "cert.pem", "key.pem", nil)
	if err != nil {
		panic(err)
	}
}
