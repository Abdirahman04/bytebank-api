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
