package models_test

import (
	"testing"

	"github.com/lifeisgo/shrimptalk/models"
)

func TestORM(t *testing.T) {
	models.RunMigrate()
}
