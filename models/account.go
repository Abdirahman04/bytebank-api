package models

import "time"

type Account struct {
  AccountID string `bson:"_id,omitempty"`
  CustomerID string `bson:"customer_id"`
  AccountType string `bson:"account_type"`
  Amount float32 `bson:"amount"`
  DateOfCreation time.Time `bson:"date_of_creation"`
}


type AccountRequest struct {
  CustomerID string `json:"customer_id"`
  AccountType string `json:"account_type"`
}
