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

type AccountResponse struct {
  AccountID string `json:"accountid"`
  CustomerID string `json:"customerid"`
  AccountType string `json:"accounttype"`
  Amount float32 `json:"amount"`
  DateOfCreation time.Time `json:"dateofcreation"`
}

func NewAccount(account AccountRequest) Account {
  return Account{
    CustomerID: account.CustomerID,
    AccountType: account.AccountType,
    Amount: 0.00,
    DateOfCreation: time.Now(),
  }
}

func NewAccountResponse(account Account) AccountResponse {
  return AccountResponse{
    AccountID: account.AccountID,
    CustomerID: account.CustomerID,
    AccountType: account.AccountType,
    Amount: account.Amount,
    DateOfCreation: account.DateOfCreation,
  }
}
