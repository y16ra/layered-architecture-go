package model

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

type User struct {
	ID   UserID
	Name UserName
}

type UserID string

func NewUserID() UserID {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	return UserID(id.String())
}

type UserName string

const (
	MinUserNameLength = 3
	MaxUserNameLength = 50
)

func NewUserName(name string) (UserName, error) {
	// validate name
	name = strings.TrimSpace(name)
	if err := UserName(name).IsValid(); err != nil {
		return "", err
	}
	return UserName(name), nil
}

func (u UserName) IsValid() error {
	if u == "" {
		return errors.New("User name cannot be empty")
	}

	if len(u) < MinUserNameLength || len(u) > MaxUserNameLength {
		return errors.New("User name must be between 3 and 50 characters long")
	}
	return nil
}

func NewUser(name UserName) *User {
	if err := name.IsValid(); err != nil {
		return nil
	}
	return &User{
		ID:   NewUserID(),
		Name: name,
	}
}
