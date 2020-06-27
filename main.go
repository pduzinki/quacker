package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"quacker/controllers"
	"quacker/models"
)

func main() {
	// config
	dbCfg := LoadDatabaseConfig()

	// services
	us := models.NewUserService(dbCfg.Dialect(), dbCfg.ConnectionInfo())

	// controllers
	userC := controllers.NewUserController(us)
	// TODO hashtagC := controllers.NewHashtagC()
	// TODO quackC := controllers.NewQuackC()

	// router
	r := mux.NewRouter()
	r.HandleFunc("/", userC.GetHomepage)
	r.HandleFunc("/login", userC.GetLogin).Methods("GET")
	r.HandleFunc("/login", userC.PostLogin).Methods("POST")
	r.HandleFunc("/signup", userC.GetSignup).Methods("GET")
	r.HandleFunc("/signup", userC.PostSignup).Methods("POST")

	// TODO later
	// /home
	// /explore
	// /{user:[a-zA-Z0-9_-]}

	http.ListenAndServe(":3000", r)
}
