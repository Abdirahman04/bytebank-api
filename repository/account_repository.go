package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/Abdirahman04/bytebank-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveAccount(account models.Account) (string, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("accounts")
  res, err := collection.InsertOne(context.Background(), account)
  if err != nil {
    log.Fatal(err)
    return "", err
  }
  return fmt.Sprintf("New account added: %v",res.InsertedID), nil
}

func GetAccounts() ([]models.Account, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("accounts")
  filter := bson.D{}
  curr, err := collection.Find(context.Background(), filter)
  if err != nil {
    log.Fatal(err)
    return nil, err
  }
  defer curr.Close(context.Background())
  var accounts []models.Account
  for curr.Next(context.Background()) {
    var account models.Account
    err := curr.Decode(&account)
    if err != nil {
      log.Fatal(err)
      return nil, err
    }
    accounts = append(accounts, account)
  }
  return accounts, nil
}

func GetAccountById(id string) (models.Account, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("accounts")
  filter := bson.M{"_id": id}
  var account models.Account
  err := collection.FindOne(context.Background(), filter).Decode(&account)
  if err != nil {
    return account, err
  }
  return account, err
}

func ChangeAmount(id string, amount float32) (string, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("accounts")
  account, err := GetAccountById(id)
  if err != nil {
    log.Fatal(err)
    return "", err
  }
  newAmount := account.Amount + amount
  filter := bson.M{"_id": id}
  update := bson.M{"$set": bson.M{"amount": newAmount}}
  res, err := collection.UpdateOne(context.Background(), filter, update)
  if err != nil {
    log.Fatal(err)
    return "", err
  }
  return fmt.Sprint("Updated", res), nil
}

func DeleteAccount(id string) (string, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("accounts")
  filter := bson.M{"_id": id}
  res, err := collection.DeleteOne(context.Background(), filter)
  if err != nil {
    log.Fatal(err)
    return "", err
  }
  return fmt.Sprint("Deleted", res.DeletedCount), nil
}