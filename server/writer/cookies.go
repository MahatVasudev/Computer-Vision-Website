package writer

import (
	"net/http"
	"time"
)

func WriteCookie(
	w http.ResponseWriter,
	key string,
	value string,
	path string,
	httpOnly bool,
	expiry time.Time,
) {
	cookie := http.Cookie{
		Name:     key,
		Value:    value,
		Path:     path,
		Expires:  expiry,
		HttpOnly: httpOnly,
	}

	http.SetCookie(w, &cookie)

}

func WriteCookieSSID(w http.ResponseWriter, value string) {
	WriteCookie(
		w,
		"SSID",
		value,
		"/",
		false,
		time.Date(time.Now().Year()+1, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
	)
}

func DeleteCookie(
	w http.ResponseWriter,
	key string,
) {
	cookie := http.Cookie{
		Name:    key,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}

	http.SetCookie(w, &cookie)
}
