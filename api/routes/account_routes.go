package routes

import (
	"net/http"

	"github.com/Abdirahman04/bytebank-api/api/handlers"
	"github.com/gorilla/mux"
)

func SetAccountRoutes(router *mux.Router) {
  router.HandleFunc("/accounts", handlers.GetAccounts).Methods(http.MethodGet)
  router.HandleFunc("/accounts/{id}", handlers.GetAccountById).Methods(http.MethodGet)
  router.HandleFunc("/accounts", handlers.PostAccount).Methods(http.MethodPost)
  router.HandleFunc("/accounts/{id}", handlers.DeleteAccount).Methods(http.MethodDelete)
}
