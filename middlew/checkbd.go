package middlew

import (
	"net/http"

	"github.com/manuelnelson7/twittit/bd"
)

func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Lost connection with the database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
