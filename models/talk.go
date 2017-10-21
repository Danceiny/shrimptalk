package models

import (
	"encoding/json"

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

type CommentDetail struct {
	NickName string
	Comment  string
}

func NewTalk() *Talk {
	talk := new(Talk)
	talk.Max = 0
	talk.TalkNameHex = GenerateHexID()
	return talk
}

func (t *Talk) AddComment(name string, comment string) {
	t.Max += 1
	t.DetailIndex = AddIndex(t.DetailIndex, t.Max, name)
	t.CommentDetail = AddDetail(t.CommentDetail, name, comment)
}

func AddIndex(detail string, max int, name string) string {
	di := map[int]string{}
	err := json.Unmarshal([]byte(detail), &di)
	max -= 1
	if err == nil {
		di[max] = name
	} else {
		di = map[int]string{max: name}
	}
	v, _ := json.Marshal(di)
	return string(v)
}

func AddDetail(detail, name, comment string) string {
	di := map[string]string{}
	err := json.Unmarshal([]byte(detail), &di)
	if err == nil {
		di[name] = comment
	} else {
		di = map[string]string{name: comment}
	}
	v, _ := json.Marshal(di)
	return string(v)
}

func (t *Talk) ToComment() []CommentDetail {
	di := map[int]string{}
	cd := map[string]string{}
	json.Unmarshal([]byte(t.DetailIndex), &di)
	json.Unmarshal([]byte(t.CommentDetail), &cd)
	cdArr := []CommentDetail{}
	for i := 0; i < t.Max; i++ {
		if name, b := di[i]; b {
			if detail, b := cd[name]; b {
				cdArr = append(cdArr, CommentDetail{name, detail})
			}
		}
	}
	return cdArr
}
