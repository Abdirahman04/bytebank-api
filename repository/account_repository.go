package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Abdirahman04/bytebank-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveAccount(account models.Account) (string, error) {
  log.Println("SaveAccount in repository hit")

  client := Connect()
  collection := client.Database("bytebank").Collection("accounts")

  res, err := collection.InsertOne(context.Background(), account)
  if err != nil {
    log.Println("error inserting account", err)
    return "", errors.New("Error adding account")
  }

  log.Println("SaveAccount in repository success")
  return fmt.Sprintf("New account added: %v",res.InsertedID), nil
}

func GetAccounts() ([]models.Account, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("accounts")
  filter := bson.D{}

  curr, err := collection.Find(context.Background(), filter)
  if err != nil {
    return nil, errors.New("no account found")
  }

  defer curr.Close(context.Background())

  var accounts []models.Account
  for curr.Next(context.Background()) {
    var account models.Account
    err := curr.Decode(&account)

    if err != nil {
      return nil, errors.New("no account found")
    }
    accounts = append(accounts, account)
  }
  return accounts, nil
}

func GetAccountById(id string) (models.Account, error) {
  objectId, err := primitive.ObjectIDFromHex(id)

  if err != nil {
    return models.Account{}, err
  }

  client := Connect()
  collection := client.Database("bytebank").Collection("accounts")
  filter := bson.M{"_id": objectId}

  var account models.Account
  err = collection.FindOne(context.Background(), filter).Decode(&account)
  if err != nil {
    return models.Account{}, errors.New("no account found")
  }

  return account, err
}

func GetAccountsByCustomerId(id string) ([]models.Account, error) {
  log.Println("GetAccountsByCustomerId rep hit")

  client := Connect()
  collection := client.Database("bytebank").Collection("accounts")
  filter := bson.M{"customer_id": id}

  curr, err := collection.Find(context.Background(), filter)
  if err != nil {
    log.Println(err.Error())
    return nil, errors.New("no account found")
  }

  defer curr.Close(context.Background())

  var accounts []models.Account
  for curr.Next(context.Background()) {
    var account models.Account
    err = curr.Decode(&account)
    if err != nil {
      log.Println(err.Error())
      continue
    }
    accounts = append(accounts, account)
  }

  return accounts, nil
}

func ChangeAmount(id string, amount float32) (string, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("accounts")

  account, err := GetAccountById(id)
  if err != nil {
    log.Println("Error getting account:", err)
    return "", err
  }

  objectId, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return "", err
  }

  newAmount := account.Amount + amount
  filter := bson.M{"_id": objectId}
  update := bson.M{"$set": bson.M{"amount": newAmount}}

  res, err := collection.UpdateOne(context.Background(), filter, update)
  if err != nil {
    log.Println("Error changing amount:", err)
    return "", errors.New("error changing amount")
  }

  return fmt.Sprint("Updated", res), nil
}

func DeleteAccount(id string) (string, error) {
  objectId, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return "", err
  }

  client := Connect()
  collection := client.Database("bytebank").Collection("accounts")
  filter := bson.M{"_id": objectId}

  res, err := collection.DeleteOne(context.Background(), filter)
  if err != nil {
    log.Fatal(err)
    return "", errors.New("error deleting account")
  }

  return fmt.Sprint("Deleted", res.DeletedCount), nil
}
