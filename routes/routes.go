// Package routes consist of router path used for handling incoming request //
package routes

import (
	"github.com/gorilla/mux"
	"gitlab.com/my-app/controllers"
)

// Route is
type Route struct{}

// Init is
func (r *Route) Init() *mux.Router {
	// Initialize controller //
	healthCheckController := controllers.InitHealthCheckController()
	playerController := controllers.InitPlayerController()
	usersController := controllers.InitUsersController()
	productController := controllers.InitProductController()

	// Initialize router //
	router := mux.NewRouter().StrictSlash(false)
	v1 := router.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/healthcheck", healthCheckController.HealthCheck).Methods("GET")
	v1.HandleFunc("/player", playerController.StorePlayer).Methods("POST")

	//User
	v1.HandleFunc("/users/store", usersController.StoreUser).Methods("POST")
	v1.HandleFunc("/users", usersController.GetUsers).Methods("GET")
	v1.HandleFunc("/users/detail/{id}", usersController.GetUserByID).Methods("GET")

	//product
	v1.HandleFunc("/product", productController.ProductList).Methods("GET")
	v1.HandleFunc("/product/detail/{id}", productController.GetProductByID).Methods("GET")
	v1.HandleFunc("/product/store", productController.StoreProduct).Methods("POST")

	return v1
}
