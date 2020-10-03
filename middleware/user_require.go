package middleware

import (
	"log"
	"net/http"

	"quacker/context"
	"quacker/models"
	"quacker/redirect"
)

// UserRequire is a middleware that verifies that there is a user logged in,
// and if not, it redirects to "/login" page
type UserRequire struct {
	models.UserService
}

// ApplyFn applies middleware to given function
func (mw *UserRequire) ApplyFn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("remember_token")
		if err != nil {
			redirect.RememberURL(w, r.URL.Path)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		user, err := mw.FindByRememberToken(cookie.Value)
		if err != nil {
			redirect.RememberURL(w, r.URL.Path)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		ctx := r.Context()
		ctx = context.WithUser(ctx, user)
		r = r.WithContext(ctx)

		log.Println("User found:", user)
		next(w, r)
	})
}
