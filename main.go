package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"quacker/controllers"
	"quacker/models"
)

func main() {
	// config
	cfg := LoadConfig()
	dbCfg := cfg.DbConfig

	// services
	services := models.NewServices(dbCfg.Dialect(), dbCfg.ConnectionInfo(),
		cfg.PasswordPepper, cfg.HmacKey)
	defer services.Close()
	// services.RebuildDatabase()
	err := services.AutoMigrate()
	if err != nil {
		panic(err)
	}

	// controllers
	userC := controllers.NewUserController(services.Us)
	// TODO hashtagC := controllers.NewHashtagC()
	// TODO quackC := controllers.NewQuackC()

	// router
	r := mux.NewRouter()
	r.HandleFunc("/", userC.GetHomepage)
	r.HandleFunc("/login", userC.GetLogin).Methods("GET")
	r.HandleFunc("/login", userC.PostLogin).Methods("POST")
	r.HandleFunc("/signup", userC.GetSignup).Methods("GET")
	r.HandleFunc("/signup", userC.PostSignup).Methods("POST")
	r.HandleFunc("/{user:[a-zA-Z0-9_-]+}", userC.GetUser).Methods("GET")
	r.HandleFunc("/quack", userC.GetNewQuack).Methods("GET")
	r.HandleFunc("/quack", userC.PostNewQuack).Methods("POST")

	// TODO
	// add 'new quack' functionality
	// add quacks to user profile

	http.ListenAndServe(":3000", r)
}
