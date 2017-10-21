package models

import (
	"encoding/json"

	"log"

	"github.com/satori/go.uuid"
)

type Talk struct {
	Base
	TalkNameHex   string
	Now           uuid.UUID `gorm:"primary_key;type:varchar(36)"`
	Max           int
	DetailIndex   string `gorm:"type:varchar(2048)"`
	CommentDetail string `gorm:"type:varchar(4096)"`
}

//type CommentIndex struct {
//	Index    int
//	NickName string
//}

func NewTalk() *Talk {
	talk := new(Talk)
	talk.Max = 0
	talk.TalkNameHex = GenerateHexID()
	return talk
}

func (t *Talk) AddComment(name string, comment string) string {
	t.Max += 1
	return Add(t.DetailIndex, name, comment)
}

func Add(detail, name, comment string) string {

	di := map[int]string{}
	err := json.Unmarshal([]byte(detail), &di)
	if err != nil {
		log.Println(err)
	}
	v, _ := json.Marshal(di)
}
