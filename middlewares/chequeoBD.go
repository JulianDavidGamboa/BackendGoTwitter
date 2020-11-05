package middlewares

import (
	"net/http"

	"github.com/JulianDavidGamboa/BackendGoTwitter/database"
)

/*ChequeoBD middleware que permite verificar la conexion con la base de datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos ", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
