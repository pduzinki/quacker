package main

import (
	"log"
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
	services := models.NewServices(dbCfg.GetDialect(), dbCfg.GetConnectionInfo(),
		cfg.PasswordPepper, cfg.HmacKey)
	defer services.Close()
	// services.RebuildDatabase()
	err := services.AutoMigrate()
	if err != nil {
		panic(err)
	}

	// controllers
	userC := controllers.NewUserController(services.Us)
	quackC := controllers.NewQuackController(services.Qs, services.Us, services.Fs, services.Hs)
	followC := controllers.NewFollowController(services.Fs, services.Us)
	hashtagC := controllers.NewHashtagController(services.Qs, services.Hs)

	// middleware
	userRequireMw := middleware.UserRequire{
		UserService: services.Us,
	}
	userLoggedMw := middleware.UserLogged{
		UserService: services.Us,
	}

	// router
	r := mux.NewRouter()
	r.HandleFunc("/", userC.GetWelcome)
	r.HandleFunc("/login", userC.GetLogin).Methods("GET")
	r.HandleFunc("/login", userC.PostLogin).Methods("POST")
	r.HandleFunc("/signup", userC.GetSignup).Methods("GET")
	r.HandleFunc("/signup", userC.PostSignup).Methods("POST")
	r.HandleFunc("/logout", userC.PostLogout).Methods("POST")
	r.HandleFunc("/home", userRequireMw.ApplyFn(quackC.GetHome)).Methods("GET")
	r.HandleFunc("/home", userRequireMw.ApplyFn(quackC.NewQuack)).Methods("POST")
	r.HandleFunc("/cookietest", userC.CookieTest).Methods("GET")

	r.HandleFunc("/hashtags/{hashtag:[a-zA-Z0-9_]+}", hashtagC.ShowQuacksByHashtag).Methods("GET")

	r.HandleFunc("/{user:[a-zA-Z0-9_-]+}/follow", userRequireMw.ApplyFn(followC.FollowUser)).Methods("POST")
	r.HandleFunc("/{user:[a-zA-Z0-9_-]+}/unfollow", userRequireMw.ApplyFn(followC.UnfollowUser)).Methods("POST")
	r.HandleFunc("/{user:[a-zA-Z0-9_-]+}/following", userRequireMw.ApplyFn(followC.Following)).Methods("GET")
	r.HandleFunc("/{user:[a-zA-Z0-9_-]+}/followers", userRequireMw.ApplyFn(followC.Followers)).Methods("GET")

	r.HandleFunc("/{user:[a-zA-Z0-9_-]+}/quack/{id:[0-9]+}", quackC.GetQuack).Methods("GET")
	r.HandleFunc("/{user:[a-zA-Z0-9_-]+}/quack/{id:[0-9]+}/delete", userRequireMw.ApplyFn(quackC.DeleteQuack)).Methods("POST")
	r.HandleFunc("/{user:[a-zA-Z0-9_-]+}", quackC.GetProfile).Methods("GET")

	// styles
	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles/"))))

	log.Println("Quacker is now working on port 3000")
	http.ListenAndServe(":3000", userLoggedMw.Apply(r))
}
