package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DB         string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	LoadServer(cfg)
	LoadData(cfg)
}

func LoadServer(iniFile *ini.File) {
	AppMode = iniFile.Section("").Key("app_mode").String()
	HttpPort = iniFile.Section("server").Key("http_port").String()
	JwtKey = iniFile.Section("server").Key("jwt_key").String()
}

func LoadData(iniFile *ini.File) {
	dbSection := iniFile.Section("database")
	DB = dbSection.Key("db").String()
	DBHost = dbSection.Key("db_host").String()
	DBPort = dbSection.Key("db_port").String()
	DBName = dbSection.Key("db_name").String()
	DBUser = dbSection.Key("db_user").String()
	DBPassword = dbSection.Key("db_password").String()
}
