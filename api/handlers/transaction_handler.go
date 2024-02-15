package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/services"
)

func PostTransaction(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var transaction models.TransactionRequest
  err := json.NewDecoder(r.Body).Decode(&transaction)
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
  }
  res, err := services.SaveTransaction(transaction)
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
  }
  json.NewEncoder(w).Encode(res)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  res, err := services.GetTransactions()
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
  }
  json.NewEncoder(w).Encode(res)
}
