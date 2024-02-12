package services

import (
	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/repository"
)

func SaveUser(user models.UserRequest) (string, error) {
  res, err := repository.SaveUser(user)
  return res, err
}

func GetUsers() ([]models.UserResponse, error) {
  res, err := repository.GetUsers()
  return res, err
}

func GetUserByEmail(email string) (models.UserResponse, error) {
  res, err := repository.GetUserByEmail(email)
  return res, err
}

func UpdateUser(email string, user models.UserRequest) error {
  return repository.UpdateUser(email, user)
}

func DeleteUser(email string) (string, error) {
  res, err := repository.DeleteUser(email)
  return res, err
}
