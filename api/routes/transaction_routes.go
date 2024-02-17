package routes

import (
	"net/http"

	"github.com/Abdirahman04/bytebank-api/api/handlers"
	"github.com/gorilla/mux"
)

func SetTransactionRoutes(router *mux.Router) {
  router.HandleFunc("/transactions", handlers.GetTransactions).Methods(http.MethodGet)
  router.HandleFunc("/transactions/{id}", handlers.GetTransactionById).Methods(http.MethodGet)
  router.HandleFunc("/transactions/account/{id}", handlers.GetTransactionsByAccountId).Methods(http.MethodGet)
  router.HandleFunc("/transactions", handlers.PostTransaction).Methods(http.MethodPost)
  router.HandleFunc("/transactions/{id}", handlers.DeleteTransaction).Methods(http.MethodDelete)
}
