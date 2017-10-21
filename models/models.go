package models

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"primary_key;type:varchar(36)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type DB = gorm.DB

var (
	db             *DB
	globalSessions *session.Manager
	migrate        = []interface{}{}
)

func init() {
	db = CreateDB()
}

func (b *Base) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV1())
	return nil
}

func connectString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%sloc=%s&charset=utf8&parseTime=true",
		"root", "", "localhost", "4000", "shrimp", "?", url.QueryEscape("Asia/Shanghai"))
}

func CreateDB() *DB {
	db, e := gorm.Open("mysql", connectString())
	if e != nil {
		panic(e)
	}
	return db

}

func ORM() *DB {
	if db == nil || db.DB().Ping() != nil {
		db = CreateDB()
	}
	db.LogMode(true)
	return db
}

func Session() *session.Manager {
	if globalSessions == nil {
		globalSessions = NewSession()
	}
	return globalSessions
}

func SetMigrate(table interface{}) {
	migrate = append(migrate, table)

}

func RunMigrate() {
	for _, v := range migrate {
		ORM().AutoMigrate(v)
	}
}

func (b Base) IsNil() bool {
	return uuid.Equal(b.ID, uuid.Nil)
}

func GenerateHexID() string {
	s := uuid.NewV1().String()
	s = strings.Replace(s, "-", "", -1)
	run := ([]rune)(s)[:32]
	return string(run)
}

func NewSession() *session.Manager {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "SessionID",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
	return globalSessions
}
