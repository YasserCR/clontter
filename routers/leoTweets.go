package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YasserCR/clontter/bd"
)

func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar la página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Pagina debe ser mayor a 0", http.StatusBadRequest)
		return
	}
	pag := int64(pagina)
	respuesta, correcto := bd.LeoTweets(ID, pag)
	if correcto == false {
		http.Error(w, "Error al leer los Tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
