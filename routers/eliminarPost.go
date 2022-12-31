package routers

import (
	"net/http"

	"github.com/DanMurguia/TT_FreeJobs/bd"
)

/*EliminarPost permite borrar un Post determinado */
func EliminarPost(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroPost(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el Post "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
