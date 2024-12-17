package routes

import (
	"fmt"
	"go-auth/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, authController *controllers.AuthController) {
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("HelloWorld")
	})
	r.HandleFunc("/register", authController.Register).Methods("POST")
}
