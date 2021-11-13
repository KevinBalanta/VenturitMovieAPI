package libs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Owner string

type DBConfig struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
}

func (c *DBConfig) InitDB() *gorm.DB {

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}

	mysqlDB, _ := db.DB()
	mysqlDB.SetMaxOpenConns(c.MaxOpenConns)
	mysqlDB.SetMaxIdleConns(c.MaxIdleConns)

	return db

}
