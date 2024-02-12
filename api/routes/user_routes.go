package routes

import (
	"net/http"

	"github.com/Abdirahman04/bytebank-api/api/handlers"
	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) {
  router.HandleFunc("/users", handlers.PostUser).Methods(http.MethodPost)
  router.HandleFunc("/users", handlers.GetUsers).Methods(http.MethodGet)
  router.HandleFunc("/users/{email}", handlers.GetUserByEmail).Methods(http.MethodGet)
  router.HandleFunc("/users/{email}", handlers.UpdateUser).Methods(http.MethodPut)
}
