package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "welcome to quacker! #quack")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)

	http.ListenAndServe(":3000", r)
}
