package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DanMurguia/TT_FreeJobs/bd"
)

/*LeoPostsSeguidores lee los posts de todos nuestros seguidores */
func LeoPostsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoPostsSeguidores(IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "Error al leer los posts", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
