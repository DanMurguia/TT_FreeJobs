package bd

import (
	"github.com/DanMurguia/TT_FreeJobs/models"
)

/*IntentoLogin realiza el chequeo de login a la BD */
func IntentoReestablecer(email string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usu, false
	}
	return usu, true
}
