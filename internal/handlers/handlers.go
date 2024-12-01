package handlers

import (
	"authService/internal/controllers"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		w.Write([]byte("Method - " + r.Method + " not allowed :("))
	} else {
		controllers.AccessMethod(w, r)
	}
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		w.Write([]byte("Method - " + r.Method + " not allowed :("))
	} else {
		controllers.RefreshMethod(w, r)
	}
}
