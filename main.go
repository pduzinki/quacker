package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"quacker/controllers"
	"quacker/middleware"
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
	quackC := controllers.NewQuackController(services.Qs, services.Us)
	// TODO hashtagC := controllers.NewHashtagC()

	// middleware
	userRequireMw := middleware.UserRequire{
		UserService: services.Us,
	}

	// router
	r := mux.NewRouter()
	r.HandleFunc("/", userC.GetWelcome)
	r.HandleFunc("/login", userC.GetLogin).Methods("GET")
	r.HandleFunc("/login", userC.PostLogin).Methods("POST")
	r.HandleFunc("/signup", userC.GetSignup).Methods("GET")
	r.HandleFunc("/signup", userC.PostSignup).Methods("POST")
	r.HandleFunc("/home", userRequireMw.ApplyFn(quackC.GetHome)).Methods("GET")
	r.HandleFunc("/home", userRequireMw.ApplyFn(quackC.NewQuack)).Methods("POST")
	r.HandleFunc("/cookietest", userC.CookieTest).Methods("GET")
	r.HandleFunc("/{user:[a-zA-Z0-9_-]+}", userC.GetUser).Methods("GET")

	http.ListenAndServe(":3000", r)
}
