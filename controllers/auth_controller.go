package controllers

import (
	"encoding/json"
	"go-auth/models"
	"go-auth/services"
	"net/http"
)

type AuthController struct {
	Service *services.AuthService
}

func (c *AuthController) Register (w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := c.Service.Register(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}