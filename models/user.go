package models

import (
	"time"
)

type User struct {
	Base
	NickNameHex string
}

func init() {
	SetMigrate(User{})
}

func NewUser() *User {
	user := new(User)
	user.NickNameHex = GenerateHexID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return user
}
