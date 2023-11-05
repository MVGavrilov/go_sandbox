package controller

import (
	"fmt"
	"mysrvr/app/model"
	"net/http"
)

func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	name := r.FormValue("name")
	lastName := r.FormValue("last_name")
	err := model.UserCreate(name, lastName, email, password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User created for "+email+" "+name+" "+lastName+"!")
}