package models

import (
	"fmt"

	"net/url"

	"github.com/lifeisgo/shrimptalk/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DB = gorm.DB

var (
	db      *DB
	migrate = []interface{}{}
)

func init() {
	db = CreateDB()
}

func connectString() string {

	host := common.GetConfig("mysql::host")
	port := common.GetConfig("mysql::port")
	user := common.GetConfig("mysql::user")
	password := common.GetConfig("mysql::password")
	database := common.GetConfig("mysql::database")
	param := "?"
	loc := common.GetConfig("mysql::loc")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%sloc=%s&charset=utf8&parseTime=true",
		user, password, host, port, database, param, url.QueryEscape("Asia/Shanghai"))
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

func SetMigrate(table interface{}) {
	migrate = append(migrate, table)

}

func RunMigrate() {
	for _, v := range migrate {
		ORM().AutoMigrate(v)
	}
}
