package models

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"strings"

	"github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type RandomKind int

const (
	T_RAND_NUM      RandomKind = iota // 纯数字
	T_RAND_LOWER                      // 小写字母
	T_RAND_UPPER                      // 大写字母
	T_RAND_LOWERNUM                   // 数字、小写字母
	T_RAND_UPPERNUM                   // 数字、大写字母
	T_RAND_ALL                        // 数字、大小写字母
)

var (
	RandomString = map[RandomKind]string{
		T_RAND_NUM:      "0123456789",
		T_RAND_LOWER:    "abcdefghijklmnopqrstuvwxyz",
		T_RAND_UPPER:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		T_RAND_LOWERNUM: "0123456789abcdefghijklmnopqrstuvwxyz",
		T_RAND_UPPERNUM: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		T_RAND_ALL:      "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
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
	s := strings.Replace(uuid.NewV1().String(), "-", "", -1)
	run := ([]rune)(s)[:32]
	return string(run)
}

func GenerateRandomString(size int, kind ...RandomKind) string {

	bytes := RandomString[T_RAND_ALL]
	if kind != nil {
		if k, b := RandomString[kind[0]]; b == true {
			bytes = k
		}
	}
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
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
