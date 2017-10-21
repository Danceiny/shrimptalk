package models_test

import (
	"testing"

	"log"

	"github.com/lifeisgo/shrimptalk/models"
	"github.com/satori/go.uuid"
)

func TestAddDetail(t *testing.T) {
	log.Println(models.AddDetail(`{"cc":"hl"}`, 3, "ee"))
}

func TestAddIndex(t *testing.T) {
	log.Println(models.AddIndex(`{"1":"dd"}`, 3, "dd"))
}

func TestTalk_AddComment(t *testing.T) {

	//tk := models.NewTalk()
	tk := new(models.Talk)
	tk2 := new(models.Talk)
	tk.TalkNameHex = "5a61c55ab62a11e78494da1584770d13"
	models.ORM().Where(tk).FirstOrInit(tk2)
	log.Println(tk2)
	tk2.AddComment("hello", "world")
	tk2.Now = uuid.FromStringOrNil("7a4cc709-b629-11e7-b1a0-da1584770d13")
	models.ORM().Where(tk).Save(tk2)

}

func TestTalk_ToComment(t *testing.T) {
	tk := new(models.Talk)
	tk2 := new(models.Talk)
	tk.TalkNameHex = "5a61c55ab62a11e78494da1584770d13"
	models.ORM().Where(tk).FirstOrInit(tk2)
	log.Println(tk2.ToComment())
}
