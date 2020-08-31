package controllers

import "net/http"

type userController struct{}

func (uc userController) serveHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Heelo from the user controller"))
}
