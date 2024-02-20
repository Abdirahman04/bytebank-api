package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Abdirahman04/bytebank-api/logger"
	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/services"
	"github.com/gorilla/mux"
)

func PostTransaction(w http.ResponseWriter, r *http.Request) {
  al := logger.NewAggregatedLogger()
  al.Info(logger.EndpointLog("PostTransaction", r))
  w.Header().Set("Content-Type", "application/json")
  var transaction models.TransactionRequest
  err := json.NewDecoder(r.Body).Decode(&transaction)
  if err != nil {
    al.Error("unable to decode transaction:", err)
    log.Println("Error decoding transaction:", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }
  res, err := services.SaveTransaction(transaction)
  if err != nil {
    al.Warn(err.Error())
    log.Println("Error saving transaction:", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }
  json.NewEncoder(w).Encode(res)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
  al := logger.NewAggregatedLogger()
  al.Info(logger.EndpointLog("GetTransactions", r))
  w.Header().Set("Content-Type", "application/json")
  res, err := services.GetTransactions()
  if err != nil {
    al.Warn(err.Error())
    log.Println("Error getting transactions:", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }
  json.NewEncoder(w).Encode(res)
}

func GetTransactionById(w http.ResponseWriter, r *http.Request) {
  al := logger.NewAggregatedLogger()
  al.Info(logger.EndpointLog("GetTransactionById", r))
  w.Header().Set("Content-Type", "application/json")
  id := mux.Vars(r)["id"]
  res, err := services.GetTransactionById(id)
  if err != nil {
    al.Warn(err.Error())
    log.Println("Error getting transaction:", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }
  json.NewEncoder(w).Encode(res)
}

func GetTransactionsByAccountId(w http.ResponseWriter, r *http.Request) {
  al := logger.NewAggregatedLogger()
  al.Info(logger.EndpointLog("GetTransactionsByAccountId", r))
  w.Header().Set("Content-Type", "application/json")
  id := mux.Vars(r)["id"]
  res, err := services.GetTransactionsByAccountId(id)
  if err != nil {
    al.Warn(err.Error())
    log.Println("Error getting transaction:", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }
  json.NewEncoder(w).Encode(res)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
  al := logger.NewAggregatedLogger()
  al.Info(logger.EndpointLog("DeleteTransaction", r))
  w.Header().Set("Content-Type", "application/json")
  id := mux.Vars(r)["id"]
  res, err := services.DeleteTransaction(id)
  if err != nil {
    al.Warn(err.Error())
    log.Println("Error deleting transaction:", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }
  fmt.Println("hand-hit")
  json.NewEncoder(w).Encode(res)
}
