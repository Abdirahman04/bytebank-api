package services

import (
	"fmt"

	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/repository"
	"github.com/Abdirahman04/bytebank-api/validations"
)

func SaveTransaction(transaction models.TransactionRequest) (string, error) {
  err := validations.ValidateTransaction(transaction)
  if err != nil {
    return "", err
  }
  _, err = repository.GetAccountById(transaction.AccountID)
  if err != nil {
    return "", err
  }
  newTransaction := models.NewTransaction(transaction)
  res, err := repository.SaveTransaction(newTransaction)
  return res, err
}

func GetTransactions() ([]models.TransactionResponse, error) {
  var transactions []models.TransactionResponse
  res, err := repository.GetTransactions()
  for _, trans := range res {
    newTransaction := models.NewTransactionResponse(trans)
    transactions = append(transactions, newTransaction)
  }
  return transactions, err
}

func GetTransactionById(id string) (models.TransactionResponse, error) {
  transaction, err := repository.GetTransactionById(id)
  newTransaction := models.NewTransactionResponse(transaction)
  return newTransaction, err
}

func GetTransactionsByAccountId(id string) ([]models.TransactionResponse, error) {
  var transactions []models.TransactionResponse
  res, err := repository.GetTransactionsByAccountId(id)
  for _, trans := range res {
    newTransaction := models.NewTransactionResponse(trans)
    transactions = append(transactions, newTransaction)
  }
  return transactions, err
}

func DeleteTransaction(id string) (string, error) {
  res, err := repository.DeleteTransaction(id)
  fmt.Println("serv-hit")
  return res, err
}

func DeleteTransactionsByAccountId(id string) (string, error) {
  res, err := repository.DeleteTransactionsByAccountId(id)
  return res, err
}
