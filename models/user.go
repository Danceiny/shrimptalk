package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lifeisgo/shrimptalk/common"
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
	user.NickNameHex = common.GenerateHexID()
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

func RandomUser() *User {
	user := new(User)
	count := 0
	ORM().Table("users").Count(&count)
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(count)
	ORM().Table("users").Offset(randNum).First(user)
	return user
}

func Users() []User {
	users := []User{}
	ORM().Table("users").Find(&users)
	return users
}
