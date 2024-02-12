package models

import (
  "time"
)

type UserRequest struct {
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
  Email string `json:"email"`
  PhoneNumber string `json:"phonenumber"`
  Pin uint `json:"pin"`
}

type User struct {
  CustomerID string `bson:"_id,omitempty"`
  FirstName string `bson:"first_name"`
  LastName string `bson:"last_name"`
  Email string `bson:"email"`
  PhoneNumber string `bson:"phone_number"`
  Pin uint `bson:"pin"`
  DateOfCreation time.Time `bson:"date_of_creation"`
}

type UserResponse struct {
  CustomerID string `json:"customerid"`
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
  Email string `json:"email"`
  PhoneNumber string `json:"phonenumber"`
  Pin uint `json:"pin"`
  DateOfCreation time.Time `json:"dateofcreation"`
}

func NewUser(user UserRequest) User {
  return User{
    FirstName: user.FirstName,
    LastName: user.LastName,
    Email: user.Email,
    PhoneNumber: user.PhoneNumber,
    Pin: user.Pin,
    DateOfCreation: time.Now(),
  }
}

func NewUserResponse(user User) UserResponse {
  return UserResponse{
    CustomerID: user.CustomerID,
    FirstName: user.FirstName,
    LastName: user.LastName,
    Email: user.Email,
    PhoneNumber: user.PhoneNumber,
    Pin: user.Pin,
    DateOfCreation: user.DateOfCreation,
  }
}
