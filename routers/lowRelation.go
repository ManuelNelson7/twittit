package routers

import (
	"net/http"

	"github.com/manuelnelson7/twittit/bd"
	"github.com/manuelnelson7/twittit/models"
)

/*BajaRelacion realiza el borrado de la relacion entre usuarios */
func LowRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relation
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.DeleteRelation(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}