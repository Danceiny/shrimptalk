package models_test

import (
	"testing"

	"log"

	"github.com/godcong/shrimptalk/models"
)

func TestAddDetail(t *testing.T) {
	log.Println(models.AddDetail(`{"cc":"hl"}`, "dd", "ee"))
}

func TestAddIndex(t *testing.T) {
	log.Println(models.AddIndex(`{"1":"dd"}`, 3, "dd"))
}
