package models

import (
	"encoding/json"

	"strings"

	"github.com/satori/go.uuid"
)

type Talk struct {
	Base
	UserID        uuid.UUID `gorm:"type:varchar(36)"`
	TalkNameHex   string    //talk主题
	Now           uuid.UUID `gorm:"type:varchar(36)"`
	Max           int
	DetailIndex   string `gorm:"type:varchar(2048)"`
	CommentDetail string `gorm:"type:varchar(4096)"`
}

type CommentDetail struct {
	NickName string
	Comment  string
}

func init() {
	SetMigrate(Talk{})
}

func NewTalk() *Talk {
	talk := new(Talk)
	talk.Max = 0
	talk.TalkNameHex = GenerateRandomString(32)
	return talk
}

func (t *Talk) AddComment(name string, comment string) {
	t.Max += 1
	t.DetailIndex = AddIndex(t.DetailIndex, t.Max, name)
	t.CommentDetail = AddDetail(t.CommentDetail, t.Max, comment)
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

func AddDetail(detail string, max int, comment string) string {
	di := map[int]string{}
	err := json.Unmarshal([]byte(detail), &di)
	max -= 1
	if err == nil {
		di[max] = comment
	} else {
		di = map[int]string{max: comment}
	}
	v, _ := json.Marshal(di)
	return string(v)
}

func (t *Talk) ToComment() []CommentDetail {
	di := map[int]string{}
	cd := map[int]string{}
	json.Unmarshal([]byte(t.DetailIndex), &di)
	json.Unmarshal([]byte(t.CommentDetail), &cd)
	cdArr := []CommentDetail{}
	for i := 0; i < t.Max; i++ {
		if name, b := di[i]; b {
			if detail, b := cd[i]; b {
				cdArr = append(cdArr, CommentDetail{name, detail})
			}
		}
	}
	return cdArr
}

func (t *Talk) ToString() string {
	di := map[int]string{}
	cd := map[int]string{}
	json.Unmarshal([]byte(t.DetailIndex), &di)
	json.Unmarshal([]byte(t.CommentDetail), &cd)
	str := ""
	for i := 0; i < t.Max; i++ {
		if name, b := di[i]; b {
			if detail, b := cd[i]; b {
				str = str + "\n" + strings.Join([]string{name, detail}, ":")
			}
		}
	}
	return str
}

func (t *Talk) Create() error {
	return ORM().Table("talks").Create(t).Error
}

func FindTalkByHex(uuid string) *Talk {
	talk := new(Talk)
	ORM().Table("talks").Where("talk_name_hex = ?", uuid).First(talk)
	return talk

}
