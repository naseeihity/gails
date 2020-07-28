package config

import (
	"log"
	"os"

	"github.com/go-ini/ini"
	"github.com/joho/godotenv"
)

// Database type:数据库配置
type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Database string
	OpenConn int
	IdleConn int
}

//DatabaseCfg 数据库配置
var DatabaseCfg = &Database{}

//Redis type:Redis配置
type Redis struct {
	Type     string
	Host     string
	Password string
	Secret   string
	MaxIdle  int
}

//RedisCfg redis配置
var RedisCfg = &Redis{}

//Server 后端配置
type Server struct {
	Port         string
	AllowOrigins []string
}

//ServerCfg Server配置
var ServerCfg = &Server{}

var cfg *ini.File

//Init init configs
func Init() {
	devPath := "config/app.ini"
	prodPath := "config/app_prod.ini"

	// load env vars
	godotenv.Load()

	env := os.Getenv("GAILS_ENV")
	var err error
	if env == PROD {
		cfg, err = ini.Load(prodPath)
		check(err)
	} else {
		// 默认使用开发环境配置
		cfg, err = ini.Load(devPath)
		check(err)
	}

	mapToStruct("database", DatabaseCfg)
	mapToStruct("redis", RedisCfg)
	mapToStruct("server", ServerCfg)
}

func mapToStruct(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.mapToStruct when ==> %s err: %v", section, err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatalf("config.Init Failed when ==> load ini files: %v", err)
	}
}
