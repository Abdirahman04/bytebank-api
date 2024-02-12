package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Abdirahman04/bytebank-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
  ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
  if err != nil { return nil }
  return client
}

func SaveUser(userRequest models.UserRequest) (string, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("users")
  user := models.NewUser(userRequest)
  res, err := collection.InsertOne(context.Background(), user)
  if err != nil {
    log.Fatal(err)
    return "", err
  }
  txt := fmt.Sprint("User account created successfully, ID:", res.InsertedID)
  return txt, nil
}

func GetUsers() ([]models.UserResponse, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("users")
  filter := bson.D{}
  curr, err := collection.Find(context.Background(), filter)
  if err != nil {
    log.Fatal(err)
    return nil, err
  }
  defer curr.Close(context.Background())
  var users []models.UserResponse
  for curr.Next(context.Background()) {
    var rawUser models.User
    err := curr.Decode(&rawUser)
    if err != nil {
      return nil, err
    }
    user := models.NewUserResponse(rawUser)
    users = append(users, user)
  }
  return users, nil
}

func GetUserByEmail(email string) (models.UserResponse, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("users")
  filter := bson.D{{"email", email}}
  var rawUser models.User
  err := collection.FindOne(context.Background(), filter).Decode(&rawUser)
  if err != nil {
    log.Fatal(err)
    return models.UserResponse{}, err
  }
  user := models.NewUserResponse(rawUser)
  return user, nil
}

func UpdateUser(email string, userRaw models.UserResponse) error {
  client := Connect()
  collection := client.Database("bytebank").Collection("users")
  filter := bson.D{{"email", email}}
 
  user := models.NewUserFromResponse(userRaw)
  update := bson.M{"$set": user}

  _, err := collection.UpdateOne(context.Background(), filter, update)
  if err != nil {
    return err
  }
  return nil
}
