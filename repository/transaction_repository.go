package repository

import (
	"context"
	"fmt"

	"github.com/Abdirahman04/bytebank-api/models"
)

func SaveTransaction(transaction models.Transaction) (string, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("transactions")
  res, err := collection.InsertOne(context.Background(), transaction)
  if err != nil {
    return "", err
  }
  return fmt.Sprintf("added %v", res.InsertedID), nil
}
