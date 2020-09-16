package middleware

import (
	"net/http"

	"quacker/context"
	"quacker/models"
)

type UserLogged struct {
	models.UserService
}

func (mw *UserLogged) ApplyFn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("remember_token")
		if err != nil {
			next(w, r)
			return
		}

		user, err := mw.UserService.FindByRememberToken(cookie.Value)
		if err != nil {
			next(w, r)
			return
		}

		ctx := r.Context()
		ctx = context.WithUser(ctx, user)
		r = r.WithContext(ctx)
		next(w, r)
	})
}

func (mw *UserLogged) Apply(next http.Handler) http.HandlerFunc {
	return mw.ApplyFn(next.ServeHTTP)
}
