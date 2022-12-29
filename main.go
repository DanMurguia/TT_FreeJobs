package main

import (
	"log"

	"github.com/DanMurguia/TT_FreeJobs/bd"
	"github.com/DanMurguia/TT_FreeJobs/handlers")

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
