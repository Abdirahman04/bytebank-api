package services

import (
	"errors"

	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/repository"
	"github.com/Abdirahman04/bytebank-api/validations"
)

func SaveUser(user models.UserRequest) (string, error) {
  err := validations.ValidateUser(user)
  if err != nil {
    return "", err
  }
  res, err := repository.SaveUser(user)
  return res, err
}

func GetUsers() ([]models.UserResponse, error) {
  res, err := repository.GetUsers()
  return res, err
}

func GetUserById(id string) (models.UserResponse, error) {
  res, err := repository.GetUserById(id)
  if err != nil {
    return models.UserResponse{}, errors.New("no user found")
  }
  user := models.NewUserResponse(res)
  return user, nil
}

func GetUserByEmail(email string) (models.UserResponse, error) {
  res, err := repository.GetUserByEmail(email)
  return res, err
}

func UpdateUser(email string, user models.UserRequest) error {
  return repository.UpdateUser(email, user)
}

func DeleteUser(id string) (string, error) {
  accounts, err := repository.GetAccountsByCustomerId(id)
  if err != nil {
    return "", err
  }
  for _, account := range accounts {
    _, err := repository.DeleteTransactionsByAccountId(account.AccountID)
    if err != nil {
      return "", err
    }
    _, err = repository.DeleteAccount(account.AccountID)
    if err != nil {
      return "", err
    }
  }

  res, err := repository.DeleteUser(id)
  return res, err
}
