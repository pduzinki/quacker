package middleware

import (
	"fmt"
	"net/http"

	"quacker/context"
	"quacker/models"
)

// UserRequire is a middleware that verifies that there is a user logged in,
// and if not, it redirects to "/login" page
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

		ctx := r.Context()
		ctx = context.WithUser(ctx, user)
		r = r.WithContext(ctx)

		fmt.Println("User found:", user)
		next(w, r)
	})
}
