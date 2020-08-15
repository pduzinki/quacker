package middleware

import (
	"fmt"
	"net/http"

	"quacker/models"
)

// UserRequire TODO add description ...
type UserRequire struct {
	models.UserService
}

// ApplyFn TODO add description ...
func (mw *UserRequire) ApplyFn(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("remember_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		user, err := mw.FindByRememberToken(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		fmt.Println("User found:", user)
		next(w, r)
	})
}
