package services

import (
	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/repository"
)

func SaveUser(user models.UserRequest) (string, error) {
  res, err := repository.SaveUser(user)
  return res, err
}
