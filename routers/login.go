package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/YasserCR/clontter/bd"
	"github.com/YasserCR/clontter/jwt"
	"github.com/YasserCR/clontter/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email no puede estar vacio"+err.Error(), 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña inválidos", 400)
		return
	}
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Error al generar el token"+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/*Como grabar una cookie*/
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
