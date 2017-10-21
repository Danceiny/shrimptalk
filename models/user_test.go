package models_test

import (
	"testing"

	"github.com/lifeisgo/shrimptalk/models"
)

func TestNewUser(t *testing.T) {
	user := models.NewUser()
	models.ORM().Save(user)
}
