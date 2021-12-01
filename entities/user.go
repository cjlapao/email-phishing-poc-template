package entities

import (
	"errors"

	"github.com/google/uuid"
)

// User entity
type User struct {
	ID        string `json:"id" bson:"_id"`
	Email     string `json:"email" bson:"email"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
}

func NewUser(email string) (*User, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	user := User{
		ID:    uuid.NewString(),
		Email: email,
	}

	return &user, nil
}
