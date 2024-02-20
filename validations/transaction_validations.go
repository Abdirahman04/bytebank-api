package validations

import (
	"errors"

	"github.com/Abdirahman04/bytebank-api/models"
)

func ValidateTransaction(transaction models.TransactionRequest) error {
  if transaction.TransactionType == "transafer" && transaction.Target == "" {
    return errors.New("no target account id provided")
  }

  typ := transaction.TransactionType
  typs := [3]string{"deposit", "withdraw", "transfer"}

  for _, transType := range typs {
    if transType == typ {
      return nil
    }
  }

  return errors.New("Invalid transaction type")
}
