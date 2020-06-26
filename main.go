package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"quacker/controllers"
	"quacker/models"
)

func main() {
	r := mux.NewRouter()

	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
		"sslmode=disable", "localhost", 5432, "postgres", "123", "quacker_dev")

	// hashtagC := controllers.NewHashtagC()
	// quackC := controllers.NewQuackC()
	us := models.NewUserService(connectionInfo)
	userC := controllers.NewUserC(us)

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
