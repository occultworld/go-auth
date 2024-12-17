package main

import (
	"go-auth/config"
	"go-auth/controllers"
	"go-auth/routes"
	"go-auth/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := config.InitDB()
	authService := &services.AuthService{DB: db}
	authController := &controllers.AuthController{Service: authService}

	r := mux.NewRouter()
	routes.AuthRoutes(r, authController)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
