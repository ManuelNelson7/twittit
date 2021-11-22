package routers

import (
	"encoding/json"
	"net/http"

	"github.com/manuelnelson7/twittit/bd"
	"github.com/manuelnelson7/twittit/models"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error with the data"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "The user email is required", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "The password must have more than 6 characters", 400)
		return
	}

	_, findit, _ := bd.CheckAlreadyExists(t.Email)
	if findit == true {
		http.Error(w, "This email is already registered", 400)
		return
	}
	_, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "error on the register "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "error on the register, fatal ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
