package services

import (
	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/repository"
)

func PostAccount(rawAccount models.AccountRequest) (string, error) {
  account := models.NewAccount(rawAccount)
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
