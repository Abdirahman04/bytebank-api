package repository

import (
	"context"
	"fmt"

	"github.com/Abdirahman04/bytebank-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetTransactions() ([]models.Transaction, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("transactions")
  filter := bson.D{}
  res, err := collection.Find(context.Background(), filter)
  if err != nil {
    return nil, err
  }
  defer res.Close(context.Background())
  var transactions []models.Transaction
  for res.Next(context.Background()) {
    var transaction models.Transaction
    err := res.Decode(&transaction)
    if err != nil {
      return nil, err
    }
    transactions = append(transactions, transaction)
  }
  return transactions, nil
}

func GetTransactionById(id string) (models.Transaction, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("transactions")
  var transaction models.Transaction
  objectId, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return transaction, err
  }
  filter := bson.M{"_id": objectId}
  err = collection.FindOne(context.Background(), filter).Decode(&transaction)
  if err != nil {
    return transaction, err
  }
  return transaction, nil
}
