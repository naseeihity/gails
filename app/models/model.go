package models

import (
	"fmt"
	"gails/app/pkg/config"
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//Init 初始化数据库
func Init() {
	connStr := "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True"
	var err error
	database := fmt.Sprintf(connStr,
		config.DatabaseCfg.User,
		config.DatabaseCfg.Password,
		config.DatabaseCfg.Host,
		config.DatabaseCfg.Database)

	db, err = gorm.Open(config.DatabaseCfg.Type, database)
	if err != nil {
		log.Fatalf("database connect err: %v", err)
	}
	db.SingularTable(true)

	db.DB().SetMaxIdleConns(config.DatabaseCfg.IdleConn)
	db.DB().SetMaxOpenConns(config.DatabaseCfg.OpenConn)

	//migration or init by sql?
}
