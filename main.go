package main

import (
	"log"

	"github.com/JulianDavidGamboa/BackendGoTwitter/database"
	"github.com/JulianDavidGamboa/BackendGoTwitter/handlers"
)

func main() {
	if database.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.Handlers()
}
