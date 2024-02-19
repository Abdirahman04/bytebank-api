package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Abdirahman04/bytebank-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
    log.Println("Error saving user:", err)
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
    log.Println("Error getting users", err)
    return nil, errors.New("No users found")
  }
  defer curr.Close(context.Background())
  var users []models.UserResponse
  for curr.Next(context.Background()) {
    var rawUser models.User
    err := curr.Decode(&rawUser)
    if err != nil {
      continue
    }
    user := models.NewUserResponse(rawUser)
    users = append(users, user)
  }
  return users, nil
}

func GetUserByEmail(email string) (models.UserResponse, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("users")
  filter := bson.M{"email": email}
  var rawUser models.User
  err := collection.FindOne(context.Background(), filter).Decode(&rawUser)
  if err != nil {
    log.Println("Error getting user by email:", err)
    return models.UserResponse{}, err
  }
  user := models.NewUserResponse(rawUser)
  return user, nil
}

func GetUserById(id string) (models.User, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("users")
  objectId, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return models.User{}, err
  }
  filter := bson.M{"_id": objectId}
  var user models.User
  err = collection.FindOne(context.Background(), filter).Decode(&user)
  if err != nil {
    return user, err
  }
  return user, nil
}

func UpdateUser(email string, user models.UserRequest) error {
  client := Connect()
  collection := client.Database("bytebank").Collection("users")
  filter := bson.D{{"email", email}}
 
  update := bson.M{"$set": bson.M{
    "first_name": user.FirstName,
    "last_name": user.LastName,
    "email": user.Email,
    "phone_number": user.PhoneNumber,
    "pin": user.Pin,
  }}

  _, err := collection.UpdateOne(context.Background(), filter, update)
  if err != nil {
    return err
  }
  return nil
}

func DeleteUser(id string) (string, error) {
  client := Connect()
  collection := client.Database("bytebank").Collection("users")
  objectId, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return "", err
  }
  filter := bson.M{"_id": objectId}
  res, err := collection.DeleteOne(context.Background(), filter)
  if err != nil {
    return "", err
  }
  return fmt.Sprintf("Deleted %v documents\n", res), nil
}
