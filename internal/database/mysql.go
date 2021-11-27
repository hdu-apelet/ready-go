package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn  = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	host = "localhost"
	user = "root"
	pass = "admin"
	port = 3306
	db   = "readygo"
)

var (
	instance *gorm.DB
)

func Get() *gorm.DB {
	return instance
}

func init() {
	connectString := fmt.Sprintf(dsn, user, pass, host, port, db)
	gormdb, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		log.Panicf("connect to database: %v", err)
	}
	log.Printf("connect to database successful")
	instance = gormdb.Debug()

	if err := instance.AutoMigrate(&User{}); err != nil {
		log.Panicf("migration database: %v", err)
	}
	log.Printf("migration done")
}
