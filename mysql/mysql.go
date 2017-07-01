package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"errors"
	"github.com/chrisho/mosquito/helper"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = newConn()
	if err != nil {
		log.Println(err)
	}
}

func newConn() (*gorm.DB, error) {

	if db != nil {
		return db, nil
	}

	setTablePrefix()
	connection, err := gorm.Open("mysql", getDataSource())
	if err != nil {
		return nil, err
	}

	db = connection

	db.DB().SetConnMaxLifetime(30 * time.Second)

	return db, err
}

func GetConn() (*gorm.DB, error){
	if db == nil {
		return nil, errors.New("connection is not exist")
	}
	return db, nil
}

func setTablePrefix() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return helper.GetEnv("MysqlPrefix") + defaultTableName
	}
}

func getDataSource() string{
	dataSourceName := ""
	dataSourceName += helper.GetEnv("MysqlUser") + ":"
	dataSourceName += helper.GetEnv("MysqlPassword") + "@"
	dataSourceName += "tcp(" + helper.GetEnv("MysqlHost") + ")/"
	dataSourceName += helper.GetEnv("MysqlName") + "?"
	dataSourceName += helper.GetEnv("MysqlParameters")

	return dataSourceName
}