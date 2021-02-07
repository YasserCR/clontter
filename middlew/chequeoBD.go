package middlew

import (
	"net/http"

	"github.com/YasserCR/clontter/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
		}
		next.ServeHTTP(w, r)
	}
}
