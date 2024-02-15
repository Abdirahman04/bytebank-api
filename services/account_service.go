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

func GetAccountById(id string) (models.AccountResponse, error) {
  res, err := repository.GetAccountById(id)
  account := models.NewAccountResponse(*res)
  return account, err
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
