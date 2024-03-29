package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/DanMurguia/TT_FreeJobs/middlew"
	"github.com/DanMurguia/TT_FreeJobs/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/enviarEnlace", middlew.ChequeoBD(routers.Reestablecer)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")

	router.HandleFunc("/post", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboPost))).Methods("POST")
	router.HandleFunc("/leoPosts", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoPosts))).Methods("GET")
	router.HandleFunc("/eliminarPost", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarPost))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoPostsSeguidores", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoPostsSeguidores))).Methods("GET")

	router.HandleFunc("/resena", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboResena))).Methods("POST")
	router.HandleFunc("/eliminarResena", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarResena))).Methods("DELETE")
	router.HandleFunc("/leoResenas", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoResenas))).Methods("GET")

	router.HandleFunc("/subirPostImg", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirPostImg))).Methods("POST")
	router.HandleFunc("/obtenerPostImg", middlew.ChequeoBD(routers.ObtenerPostImg)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
