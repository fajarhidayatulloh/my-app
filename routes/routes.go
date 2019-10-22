// Package routes consist of router path used for handling incoming request //
package routes

import (
	"github.com/gorilla/mux"
	"github.com/my-app/controllers"
)

// Route is
type Route struct{}

// Init is
func (r *Route) Init() *mux.Router {
	// Initialize controller //
	healthCheckController := controllers.InitHealthCheckController()
	usersController := controllers.InitUsersController()

	// Initialize router //

	router := mux.NewRouter().StrictSlash(false)
	v1 := router.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/ping", healthCheckController.HealthCheck).Methods("GET")
	//User
	v1.HandleFunc("/users/store", usersController.StoreUser).Methods("POST")
	v1.HandleFunc("/users", usersController.GetUsers).Methods("GET")
	v1.HandleFunc("/users/profile/{id}", usersController.GetUserByID).Methods("GET")

	return v1
}
