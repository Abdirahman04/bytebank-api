package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/services"
	"github.com/gorilla/mux"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var newUser models.UserRequest
  err := json.NewDecoder(r.Body).Decode(&newUser)
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
  }
  res, err := services.SaveUser(newUser)
  if err != nil {
    log.Printf("Error saving user: %v", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(res)
} 

func GetUsers(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  users, err := services.GetUsers()
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  id := mux.Vars(r)["id"]
  user, err := services.GetUserById(id)
  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(err)
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(user)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  email := mux.Vars(r)["email"]
  user, err := services.GetUserByEmail(email)
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  email := mux.Vars(r)["email"]
  var user models.UserRequest
  json.NewDecoder(r.Body).Decode(&user)
  err := services.UpdateUser(email, user)
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
    fmt.Println(err)
    return
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode("Update successful")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  id := mux.Vars(r)["id"]
  res, err := services.DeleteUser(id)
  if err != nil {
    http.Error(w, "Bad request", http.StatusBadRequest)
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(res)
}
