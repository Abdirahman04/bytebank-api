package validations

import (
	"errors"
	"net/mail"
	"regexp"

	"github.com/Abdirahman04/bytebank-api/models"
)

func ValidateUser(user models.UserRequest) error {
  if len(user.FirstName) < 3 {
    return errors.New("The first name should not be less than 3 characters")
  }
  if len(user.LastName) < 3 {
    return errors.New("The last name should not be less than 3 characters")
  }
  if _, err := mail.ParseAddress(user.Email); err != nil {
    return errors.New("Invalid email address")
  }
  if !regexp.MustCompile(`^\+\d{3}-\d{3}-\d{3}-\d{3}$`).MatchString(user.PhoneNumber) {
    return errors.New("Invalid phonenumber")
  }
  if user.Pin < 1000 || user.Pin > 9999 {
    return errors.New("Pin should be 4 digits long")
  }
  return nil
}
