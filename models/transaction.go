package models

import "time"

type Transaction struct {
  TransactionID string `bson:"_id,omitempty"`
  AccountID string `bson:"account_id"`
  TransactionType string `bson:"transaction_type"`
  Target string `bson:"target"`
  Balance float32 `bson:"balance"`
  DateOfCreation time.Time `bson:"date_of_creation"`
}
