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

type TransactionRequest struct {
  AccountID string `json:"account_id"`
  TransactionType string `json:"transaction_type"`
  Target string `json:"target"`
  Balance float32 `json:"balance"`
}

func NewTransaction(transaction TransactionRequest) Transaction {
  return Transaction{
    AccountID: transaction.AccountID,
    TransactionType: transaction.TransactionType,
    Target: transaction.Target,
    Balance: transaction.Balance,
    DateOfCreation: time.Now(),
  }
}
