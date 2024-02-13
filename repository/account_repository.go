package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/Abdirahman04/bytebank-api/models"
)

func SaveAccount(accountRequest models.AccountRequest) (string, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("account")
  account := models.NewAccount(accountRequest)
  res, err := collection.InsertOne(context.Background(), account)
  if err != nil {
    log.Fatal(err)
    return "", err
  }
  return fmt.Sprintf("New account added: %v",res.InsertedID), nil
}
