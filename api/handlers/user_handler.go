package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/repository"
	"github.com/gorilla/mux"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var newUser models.UserRequest
  err := json.NewDecoder(r.Body).Decode(&newUser)
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
  }
  res, err := repository.SaveUser(newUser)
  if err != nil {
    log.Fatal(err)
    return
  }
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(res)
} 

func GetUsers(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  users, err := repository.GetUsers()
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(users)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  email := mux.Vars(r)["email"]
  user, err := repository.GetUserByEmail(email)
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  email := mux.Vars(r)["email"]
  var user models.UserResponse
  json.NewDecoder(r.Body).Decode(&user)
  err := repository.UpdateUser(email, user)
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
    return
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode("Update successful")
}
