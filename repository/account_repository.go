package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/Abdirahman04/bytebank-api/models"
	"go.mongodb.org/mongo-driver/bson"
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

func GetAccounts() ([]models.AccountResponse, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("account")
  filter := bson.D{}
  curr, err := collection.Find(context.Background(), filter)
  if err != nil {
    log.Fatal(err)
    return nil, err
  }
  defer curr.Close(context.Background())
  var accounts []models.AccountResponse
  for curr.Next(context.Background()) {
    var rawAccount models.Account
    err := curr.Decode(&rawAccount)
    if err != nil {
      log.Fatal(err)
      return nil, err
    }
    account := models.NewAccountResponse(rawAccount)
    accounts = append(accounts, account)
  }
  return accounts, nil
}
