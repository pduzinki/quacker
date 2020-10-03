package redirect

import (
	"errors"
	"net/http"
	"time"
)

// ErrCookieNotFound signals that cookie file was not found
var ErrCookieNotFound = errors.New("redirect: failed to read cookie")

// RememberURL saves the url to redirect to later
func RememberURL(w http.ResponseWriter, url string) {
	cookie := http.Cookie{
		Name:     "redirect_url",
		Value:    url,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Minute * 5),
	}

	http.SetCookie(w, &cookie)
}

// GetURL return url to redirect to next
func GetURL(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("redirect_url")
	if err != nil {
		return "", ErrCookieNotFound
	}

	url := cookie.Value
	resetURLCookie(w)

	return url, nil
}

func resetURLCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "redirect_url",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Now(),
	}

	http.SetCookie(w, &cookie)
}
