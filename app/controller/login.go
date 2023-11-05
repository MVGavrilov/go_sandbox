package controller

import (
	"fmt"
	"mysrvr/app/model"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
)

func LoginPOST(w http.ResponseWriter, r *http.Request) {
	//get email and pass, check pass
	email := r.FormValue("email")
	password := r.FormValue("password")
	usr, err := model.UserByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, err)
		return
	}
	if usr.Password != password {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Wrong password")
		return
	}
	
	key := []byte("mysecret")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": usr.ID,
		"logged_in": "yes",
		"access_level": 0,
	})
	s, _ := t.SignedString(key)
	cookie := &http.Cookie{
		Name: "access_token",
		Value: s,
		MaxAge: 60,
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Logged in as "+usr.Name+" "+usr.LastName)
}

func LogoutGET(w http.ResponseWriter, r *http.Request) {
	//logout
	key := []byte("mysecret")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": 0,
		"logged_in": "no",
		"access_level": 0,
	})
	s, _ := t.SignedString(key)
	cookie := &http.Cookie{
		Name: "access_token",
		Value: s,
		MaxAge: 60,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
