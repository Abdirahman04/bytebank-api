package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/services"
)

func PostAccount(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var account models.AccountRequest
  err := json.NewDecoder(r.Body).Decode(&account)
  if err != nil {
    log.Fatal(err)
  }
  res, err := services.PostAccount(account)
  if err != nil {
    log.Fatal(err)
  }
  json.NewEncoder(w).Encode(res)
}
