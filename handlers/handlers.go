package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JulianDavidGamboa/BackendGoTwitter/middlewares"
	"github.com/JulianDavidGamboa/BackendGoTwitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers seteo mi puerto, el Handerl y pongo a escuchar al servidor*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.VerPerfil))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
