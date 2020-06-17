package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"quacker/controllers"
)

func main() {
	r := mux.NewRouter()

	feedC := controllers.NewFeedC()
	userC := controllers.NewUserC()

	r.HandleFunc("/", feedC.GetHomepage)
	r.HandleFunc("/login", userC.GetLogin).Methods("GET")
	r.HandleFunc("/login", userC.PostLogin).Methods("POST")
	r.HandleFunc("/signup", userC.GetSignup).Methods("GET")
	r.HandleFunc("/signup", userC.PostSignup).Methods("POST")

	http.ListenAndServe(":3000", r)
}
