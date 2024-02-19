package services

import (
	"log"

	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/repository"
	"github.com/Abdirahman04/bytebank-api/validations"
)

func PostAccount(rawAccount models.AccountRequest) (string, error) {
  log.Println("PostAccount in services hit")
  _, err := GetUserById(rawAccount.CustomerID)
  if err != nil {
    return "", err
  }
  account := models.NewAccount(rawAccount)
  err = validations.ValidateAccount(account)
  if err != nil {
    return "", err
  }
  res, err := repository.SaveAccount(account)
  return res, err
}

func GetAccounts() ([]models.AccountResponse, error) {
  res, err := repository.GetAccounts()
  var accounts []models.AccountResponse
  for _, account := range res {
    newAccount := models.NewAccountResponse(account)
    accounts = append(accounts, newAccount)
  }
  return accounts, err
}

func GetAccountById(id string) (models.AccountResponse, error) {
  res, err := repository.GetAccountById(id)
  account := models.NewAccountResponse(*res)
  return account, err
}

func GetAccountsByCustomerId(id string) ([]models.AccountResponse, error) {
  log.Println("GetAccountsByCustomerId serv hit")
  var accounts []models.AccountResponse
  res, err := repository.GetAccountsByCustomerId(id)
  if err != nil {
    return nil, err
  }
  for _, account := range res {
    newAccount := models.NewAccountResponse(account)
    accounts = append(accounts, newAccount)
  }
  return accounts, nil
}

func DepositAccount(id string, amount float32) (string, error) {
  res, err := repository.ChangeAmount(id, amount)
  return res, err
}

func WithdrawAccount(id string, amount float32) (string, error) {
  res, err := repository.ChangeAmount(id, -amount)
  return res, err
}

func DeleteAccount(id string) (string, error) {
  res, err := repository.DeleteAccount(id)
  return res, err
}
