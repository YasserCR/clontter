package main

import (
	"log"

	"github.com/YasserCR/clontter/bd"
	"github.com/YasserCR/clontter/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
