package handlers

import (
	"authService/internal/controllers"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controllers.AccessMethod(w, r)
	case "POST":
		w.Write([]byte("Method - " + r.Method))
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
