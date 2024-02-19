package validations

import (
	"errors"

	"github.com/Abdirahman04/bytebank-api/models"
)

func ValidateAccount(account models.Account) error {
  typ := account.AccountType
  accountTypes := [3]string{"savings", "checking", "investment"}
  for _, t := range accountTypes {
    if t == typ {
      return nil
    }
  }
  return errors.New("Invalid account type")
}
