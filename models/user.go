package models

import (
	"fmt"
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

func FindUserByHex(id string) *User {
	user := new(User)
	ORM().Where("nick_name_hex = ?", id).Find(user)
	return user
}

func FindUser(uuid string) *User {
	user := new(User)
	ORM().Where("id = ?", uuid).Find(user)
	return user
}


func AddUser(u *User) {
	err := db.Create(u).Error
	fmt.Println("err:", err)
}
