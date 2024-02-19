package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/services"
	"github.com/gorilla/mux"
)

func PostAccount(w http.ResponseWriter, r *http.Request) {
  log.Println("PostAccount in handlers hit")
  w.Header().Set("Content-Type", "application/json")
  var account models.AccountRequest
  err := json.NewDecoder(r.Body).Decode(&account)
  if err != nil {
    log.Println("Error decoding account", err)
    w.WriteHeader(http.StatusInternalServerError)
  }
  res, err := services.PostAccount(account)
  if err != nil {
    log.Println("Error saving account: ", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }
  json.NewEncoder(w).Encode(res)
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  res, err := services.GetAccounts()
  if err != nil {
    log.Println("Error getting account:", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }
  json.NewEncoder(w).Encode(res)
}

func GetAccountById(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  id := mux.Vars(r)["id"]
  res, err := services.GetAccountById(id)
  if err != nil {
    log.Println("Error getting account by id:", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(res)
}

func GetAccountByCustomerId(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  id := mux.Vars(r)["id"]
  res, err := services.GetAccountById(id)
  if err != nil {
    log.Println("Error getting account by customer id:", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(res)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  id := mux.Vars(r)["id"]
  res, err := services.DeleteAccount(id)
  if err != nil {
    log.Fatal(err)
  }
  json.NewEncoder(w).Encode(res)
}
