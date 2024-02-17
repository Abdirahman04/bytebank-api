package routes

import (
	"net/http"

	"github.com/Abdirahman04/bytebank-api/api/handlers"
	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) {
  router.HandleFunc("/users", handlers.PostUser).Methods(http.MethodPost)
  router.HandleFunc("/users", handlers.GetUsers).Methods(http.MethodGet)
  router.HandleFunc("/users/{id}", handlers.GetUserById).Methods(http.MethodGet)
  router.HandleFunc("/users/email/{email}", handlers.GetUserByEmail).Methods(http.MethodGet)
  router.HandleFunc("/users/{email}", handlers.UpdateUser).Methods(http.MethodPut)
  router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods(http.MethodDelete)
}
