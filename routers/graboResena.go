package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DanMurguia/TT_FreeJobs/bd"
	"github.com/DanMurguia/TT_FreeJobs/models"
)

/*Graboreseña permite grabar reseñas en la base de datos */
func GraboResena(w http.ResponseWriter, r *http.Request) {
	var resena models.Resena
	err := json.NewDecoder(r.Body).Decode(&resena)

	registro := models.GraboResena{
		UserID:       IDUsuario,
		UserResenaID: resena.UserResenaID,
		Mensaje:      resena.Mensaje,
		Calificacion: resena.Calificacion,
		Fecha:        time.Now(),
	}

	_, status, err := bd.InsertoResena(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la reseña", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
