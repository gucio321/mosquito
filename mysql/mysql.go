package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type Mysql struct {
	db *gorm.DB
}

var (
	mysql *Mysql
)

func NewConn(dataSource string, logger log.Logger) (*gorm.DB, error) {

	if mysql.db != nil {
		return mysql.db, nil
	}

	db, err := gorm.Open("mysql", dataSource)
	db.SetLogger(logger)
	if err != nil {
		log.Error(err)
	}

	db.DB().SetConnMaxLifetime(30 * time.Second)

	mysql = &Mysql{
		db,
	}

	return mysql.db, err
}

func getConn() (*gorm.DB, error){
	if mysql.db == nil {
		return nil, error("not exist connection")
	}
	return mysql.db, nil
}