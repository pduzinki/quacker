package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"quacker/controllers"
)

func main() {
	r := mux.NewRouter()

	// hashtagC := controllers.NewHashtagC()
	// quackC := controllers.NewQuackC()
	userC := controllers.NewUserC()

	r.HandleFunc("/", userC.GetHomepage)
	r.HandleFunc("/login", userC.GetLogin).Methods("GET")
	r.HandleFunc("/login", userC.PostLogin).Methods("POST")
	r.HandleFunc("/signup", userC.GetSignup).Methods("GET")
	r.HandleFunc("/signup", userC.PostSignup).Methods("POST")

	// TODOs
	// /home
	// /explore
	// /{user:[a-zA-Z0-9_-]}

	http.ListenAndServe(":3000", r)
}
