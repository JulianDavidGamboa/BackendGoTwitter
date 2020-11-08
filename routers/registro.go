package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JulianDavidGamboa/BackendGoTwitter/database"
	"github.com/JulianDavidGamboa/BackendGoTwitter/models"
)

/*Registro es la funcion para crear en la BD el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres ", 400)
		return
	}

	_, encontrado, _ := database.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email ", 400)
		return
	}

	_, status, err := database.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario ", 400)
		return
	}

	log.Println("Usuario registrado correctamente")

	w.WriteHeader(http.StatusCreated)

}
