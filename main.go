package main

import (
	"TogetherAndStronger/routes/auth"
	"TogetherAndStronger/routes/db/db_handler"
	"TogetherAndStronger/routes/entity/activity"
	"TogetherAndStronger/routes/entity/category"
	"TogetherAndStronger/routes/entity/company"
	"TogetherAndStronger/routes/entity/material"
	"TogetherAndStronger/routes/entity/option"
	"TogetherAndStronger/routes/entity/place"
	"TogetherAndStronger/routes/entity/prestataire"
	"TogetherAndStronger/routes/entity/room"
	"TogetherAndStronger/routes/entity/salary"
	"TogetherAndStronger/routes/entity/team"
	"TogetherAndStronger/routes/entity/teambuilding"
	"TogetherAndStronger/routes/faq"
	"TogetherAndStronger/routes/list"
	"TogetherAndStronger/routes/lookfor"
	"TogetherAndStronger/routes/signup"
	"fmt"
	"log"
	"net/http"
	_ "strings"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Best angular website of the century incoming... ")
	fmt.Println("Got one")
}

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, req)
	log.Printf("%s, %s, %v", req.Method, req.URL.Path, time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

func main() {
	mux := http.NewServeMux()

	fmt.Println("Server starting...")

	// Start the server
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./images"))))
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/faq", faq.Faq)
	mux.HandleFunc("/db/create", db_handler.Create)
	mux.HandleFunc("/db/select", db_handler.Select)
	mux.HandleFunc("/db/update", db_handler.Update)
	mux.HandleFunc("/db/delete", db_handler.Delete)
	mux.HandleFunc("/auth/LoginCompany", auth.LoginCompany)
	mux.HandleFunc("/auth/LoginPresta", auth.LoginPresta)
	mux.HandleFunc("/auth/LoginSalary", auth.LoginSalary)
	mux.HandleFunc("/auth/LoginAdmin", auth.LoginAdmin)
	mux.HandleFunc("/signup/SignupSalary", signup.SignupSalary)
	mux.HandleFunc("/signup/SignupCompany", signup.SignupCompany)
	mux.HandleFunc("/signup/SignupPresta", signup.SignupPresta)
	mux.HandleFunc("/lookfor/LookForSalary", lookfor.LookForSalary)
	mux.HandleFunc("/lookfor/LookForCompany", lookfor.LookForCompany)
	mux.HandleFunc("/lookfor/LookForPresta", lookfor.LookForPresta)
	mux.HandleFunc("/salary/getActivities", salary.GetSalaryActivities)
	mux.HandleFunc("/salary/addInfo", salary.AddSalaryInfo)
	mux.HandleFunc("/salary/deleteInfo", salary.DeleteSalaryInfo)
	mux.HandleFunc("/salary/getInfos", salary.GetSalaryInfo)
	mux.HandleFunc("/salary/addInfos", salary.AddSalaryInfo)

	mux.HandleFunc("/list/infos", list.ListInfos)
	mux.HandleFunc("/list/category", list.ListCategory)
	mux.HandleFunc("/list/activity", list.ListActivities)

	mux.HandleFunc("/activity/addActivity", activity.AddActivity)
	mux.HandleFunc("/activity/deleteActivity", activity.DeleteActivity)
	mux.HandleFunc("/activity/getActivity", activity.GetActivity)
	mux.HandleFunc("/activity/updateActivity", activity.UpdateActivity)
	mux.HandleFunc("/activity/makeActivity", activity.MakeActivity)
	mux.HandleFunc("/activity/getCategory", activity.GetCategory)
	mux.HandleFunc("/activity/addCategory", activity.AddCategory)
	mux.HandleFunc("/activity/deleteCategory", activity.DeleteCategory)
	mux.HandleFunc("/activity/addAnime", activity.AddAnime)
	mux.HandleFunc("/activity/deleteAnime", activity.DeleteAnime)
	mux.HandleFunc("/activity/getAnime", activity.GetAnime)
	mux.HandleFunc("/activity/getActivityOption", activity.GetActivityOption)
	mux.HandleFunc("/activity/addActivityOption", activity.AddActivityOption)
	mux.HandleFunc("/activity/deleteActivityOption", activity.DeleteActivityOption)

	mux.HandleFunc("/place/addPlace", place.AddPlace)
	mux.HandleFunc("/place/deletePlace", place.DeletePlace)
	mux.HandleFunc("/place/getPlace", place.GetPlace)
	mux.HandleFunc("/place/updatePlace", place.UpdatePlace)

	mux.HandleFunc("/presta/addPresta", prestataire.AddPresta)
	mux.HandleFunc("/presta/deletePresta", prestataire.DeletePresta)
	mux.HandleFunc("/presta/getPresta", prestataire.GetPresta)
	mux.HandleFunc("/presta/updatePresta", prestataire.UpdatePresta)
	mux.HandleFunc("/presta/addAnime", prestataire.AddAnime)
	mux.HandleFunc("/presta/deleteAnime", prestataire.DeleteAnime)
	mux.HandleFunc("/presta/getAnime", prestataire.GetAnime)

	mux.HandleFunc("/company/addCompany", company.AddCompany)
	mux.HandleFunc("/company/deleteCompany", company.DeleteCompany)
	mux.HandleFunc("/company/getCompany", company.GetCompany)
	mux.HandleFunc("/company/updateCompany", company.UpdateCompany)

	mux.HandleFunc("/room/addRoom", room.AddRoom)
	mux.HandleFunc("/room/deleteRoom", room.DeleteRoom)
	mux.HandleFunc("/room/getRoom", room.GetRoom)
	mux.HandleFunc("/room/updateRoom", room.UpdateRoom)
	mux.HandleFunc("/room/getActivity", room.GetActivity)
	mux.HandleFunc("/room/addActivity", room.AddActivity)
	mux.HandleFunc("/room/deleteActivity", room.DeleteActivity)

	mux.HandleFunc("/category/add", category.Add)
	mux.HandleFunc("/category/delete", category.Delete)
	mux.HandleFunc("/category/get", category.Get)
	mux.HandleFunc("/category/update", category.Update)

	mux.HandleFunc("/teambuilding/add", teambuilding.Add)
	mux.HandleFunc("/teambuilding/delete", teambuilding.Delete)
	mux.HandleFunc("/teambuilding/get", teambuilding.Get)
	mux.HandleFunc("/teambuilding/update", teambuilding.Update)
	mux.HandleFunc("/teambuilding/addActivity", teambuilding.AddActivity)
	mux.HandleFunc("/teambuilding/getActivity", teambuilding.GetActivity)
	mux.HandleFunc("/teambuilding/getPresta", teambuilding.GetPresta)
	mux.HandleFunc("/teambuilding/getMateriel", teambuilding.GetMateriel)
	mux.HandleFunc("/teambuilding/getSalle", teambuilding.GetSalle)
	mux.HandleFunc("/teambuilding/addEngage", teambuilding.AddEngage)
	mux.HandleFunc("/teambuilding/deleteEngage", teambuilding.DeleteEngage)

	mux.HandleFunc("/team/add", team.Add)
	mux.HandleFunc("/team/delete", team.Delete)
	mux.HandleFunc("/team/get", team.Get)
	mux.HandleFunc("/team/update", team.Update)
	mux.HandleFunc("/team/getSalary", team.GetSub)

	mux.HandleFunc("/option/add", option.Add)
	mux.HandleFunc("/option/delete", option.Delete)
	mux.HandleFunc("/option/get", option.Get)
	mux.HandleFunc("/option/update", option.Update)
	mux.HandleFunc("/option/make", option.Make)

	mux.HandleFunc("/material/add", material.Add)
	mux.HandleFunc("/material/delete", material.Delete)
	mux.HandleFunc("/material/get", material.Get)
	mux.HandleFunc("/material/update", material.Update)
	mux.HandleFunc("/material/addRental", material.AddRental)
	mux.HandleFunc("/material/deleteRental", material.DeleteRental)

	//handler := cors.Default().Handler(mux)
	handler := NewLogger(mux)
	err := http.ListenAndServeTLS(":9001", "cert.pem", "key.pem", handler)
	if err != nil {
		panic(err)
	}
}
