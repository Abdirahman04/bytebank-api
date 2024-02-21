package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Abdirahman04/bytebank-api/logger"
	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/services"
	"github.com/gorilla/mux"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
  al := logger.NewAggregatedLogger()
  al.Info(logger.EndpointLog("PostUser", r))

  w.Header().Set("Content-Type", "application/json")

  var newUser models.UserRequest
  err := json.NewDecoder(r.Body).Decode(&newUser)
  if err != nil {
    al.Error("unable to decode user:", err)
    log.Println("Error decoding user:", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }

  res, err := services.SaveUser(newUser)
  if err != nil {
    al.Warn(err.Error())
    log.Println("Error saving user:", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }

  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(res)
} 

func GetUsers(w http.ResponseWriter, r *http.Request) {
  al := logger.NewAggregatedLogger()
  al.Info(logger.EndpointLog("GetUsers", r))

  w.Header().Set("Content-Type", "application/json")

  users, err := services.GetUsers()
  if err != nil {
    al.Warn(err.Error())
    log.Println("Can't get users:", err)
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(err.Error())
    return
  }

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
  al := logger.NewAggregatedLogger()
  al.Info(logger.EndpointLog("GetUserById", r))

  w.Header().Set("Content-Type", "application/json")

  id := mux.Vars(r)["id"]
  user, err := services.GetUserById(id)
  if err != nil {
    al.Warn(err.Error())
    log.Println("Error getting user:", err)
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(err.Error())
    return
  }

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(user)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
  al := logger.NewAggregatedLogger()
  al.Info(logger.EndpointLog("GetUserByEmail", r))

  w.Header().Set("Content-Type", "application/json")

  email := mux.Vars(r)["email"]
  user, err := services.GetUserByEmail(email)
  if err != nil {
    al.Warn(err.Error())
    log.Println("Error getting email:", err)
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(err.Error())
    return
  }

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
  al := logger.NewAggregatedLogger()
  al.Info(logger.EndpointLog("UpdateUser", r))

  w.Header().Set("Content-Type", "application/json")

  id := mux.Vars(r)["id"]
  var user models.UserRequest
  json.NewDecoder(r.Body).Decode(&user)
  err := services.UpdateUser(id, user)
  if err != nil {
    al.Warn(err.Error())
    w.WriteHeader(http.StatusNotFound)
    log.Println("Error updating user:", err)
    json.NewEncoder(w).Encode(err.Error())
    return
  }

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode("Update successful")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
  al := logger.NewAggregatedLogger()
  al.Info(logger.EndpointLog("DeleteUser", r))

  w.Header().Set("Content-Type", "application/json")

  id := mux.Vars(r)["id"]
  res, err := services.DeleteUser(id)
  if err != nil {
    al.Warn(err.Error())
    log.Println("Error deleting user:", err)
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(err.Error())
    return
  }

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(res)
}
