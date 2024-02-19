package services

import (
	"errors"
	"fmt"

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
    txt := fmt.Sprintf("no user with id %v found", id)
    return models.UserResponse{}, errors.New(txt)
  }
  user := models.NewUserResponse(res)
  return user, nil
}

func GetUserByEmail(email string) (models.UserResponse, error) {
  res, err := repository.GetUserByEmail(email)
  if err != nil {
    txt := fmt.Sprintf("no user with email %v found", email)
    return models.UserResponse{}, errors.New(txt)
  }
  return res, nil
}

func UpdateUser(email string, user models.UserRequest) error {
  if err := repository.UpdateUser(email, user); err != nil {
    return errors.New("Error updating user")
  }
  return nil
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
