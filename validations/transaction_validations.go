package validations

import (
	"errors"

	"github.com/Abdirahman04/bytebank-api/models"
)

func ValidateTransaction(transaction models.TransactionRequest) error {
  typ := transaction.TransactionType
  typs := [3]string{"deposit", "withdraw", "transafer"}

  for _, transType := range typs {
    if transType == typ {
      return nil
    }
  }

  return errors.New("Invalid transaction type")
}
