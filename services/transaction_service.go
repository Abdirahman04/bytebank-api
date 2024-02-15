package services

import (
	"github.com/Abdirahman04/bytebank-api/models"
	"github.com/Abdirahman04/bytebank-api/repository"
)

func SaveTransaction(transaction models.TransactionRequest) (string, error) {
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
