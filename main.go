package main

import (
	"TogetherAndStronger/routes/auth"
	"TogetherAndStronger/routes/db/db_handler"
	"TogetherAndStronger/routes/entity/activity"
	"TogetherAndStronger/routes/entity/salary"
	"TogetherAndStronger/routes/faq"
	"TogetherAndStronger/routes/list"
	"TogetherAndStronger/routes/lookfor"
	"TogetherAndStronger/routes/signup"
	"fmt"
	"net/http"
	_ "strings"

	"github.com/rs/cors"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Best angular website of the century incoming... ")
	fmt.Println("Got one")
}

func main() {

	mux := http.NewServeMux()

	fmt.Println("Server starting...")

	// Start the server
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./images"))))
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/faq", faq.Faq)
	mux.HandleFunc("/db/create", db_handler.Create)
	mux.HandleFunc("/db/select", db_handler.Select)
	mux.HandleFunc("/db/update", db_handler.Update)
	mux.HandleFunc("/db/delete", db_handler.Delete)
	mux.HandleFunc("/auth/LoginCompany", auth.LoginCompany)
	mux.HandleFunc("/auth/LoginPresta", auth.LoginPresta)
	mux.HandleFunc("/auth/LoginSalary", auth.LoginSalary)
	mux.HandleFunc("/signup/SignupSalary", signup.SignupSalary)
	mux.HandleFunc("/signup/SignupCompany", signup.SignupCompany)
	mux.HandleFunc("/signup/SignupPresta", signup.SignupPresta)
	mux.HandleFunc("/lookfor/LookForSalary", lookfor.LookForSalary)
	mux.HandleFunc("/lookfor/LookForCompany", lookfor.LookForCompany)
	//	mux.HandleFunc("/lookfor/LookForPresta", lookfor.LookForPresta)
	mux.HandleFunc("/salary/getActivities", salary.GetSalaryActivities)
	mux.HandleFunc("/salary/addInfo", salary.AddSalaryInfo)
	mux.HandleFunc("/salary/deleteInfo", salary.DeleteSalaryInfo)
	mux.HandleFunc("/salary/getInfos", salary.GetSalaryInfo)
	mux.HandleFunc("/salary/addInfos", salary.AddSalaryInfo)
	mux.HandleFunc("/list/infos", list.ListInfos)
	mux.HandleFunc("/activity/addActivity", activity.AddActivity)
	mux.HandleFunc("/activity/deleteActivity", activity.DeleteActivity)
	mux.HandleFunc("/activity/getActivity", activity.GetActivity)
	mux.HandleFunc("/activity/updateActivity", activity.UpdateActivity)
	mux.HandleFunc("/activity", list.ListActivities)

	handler := cors.Default().Handler(mux)

	err := http.ListenAndServeTLS(":9000", "cert.pem", "key.pem", handler)
	if err != nil {
		panic(err)
	}
}
