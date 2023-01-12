package routers

import (
	"io"
	"net/http"
	"os"
)

/*ObtenerPostImg envia la imagen al HTTP */
func ObtenerPostImg(w http.ResponseWriter, r *http.Request) {

	postImg := r.URL.Query().Get("postImg")
	if len(postImg) < 1 {
		http.Error(w, "Debe enviar el parámetro nombre de imagen del post", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/posts/" + postImg)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
}
