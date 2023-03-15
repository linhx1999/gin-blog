package config

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

	QiniuSever string
	AccessKey  string
	SecretKey  string
	Bucket     string
)

func init() {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	loadServer(cfg)
	loadData(cfg)
	loadQiniu(cfg)
}

func loadServer(iniFile *ini.File) {
	AppMode = iniFile.Section("").Key("app_mode").String()
	HttpPort = iniFile.Section("server").Key("http_port").String()
	JwtKey = iniFile.Section("server").Key("jwt_key").String()
}

func loadData(iniFile *ini.File) {
	dbSection := iniFile.Section("database")
	DB = dbSection.Key("db").String()
	DBHost = dbSection.Key("db_host").String()
	DBPort = dbSection.Key("db_port").String()
	DBName = dbSection.Key("db_name").String()
	DBUser = dbSection.Key("db_user").String()
	DBPassword = dbSection.Key("db_password").String()
}

func loadQiniu(iniFile *ini.File) {
	section := iniFile.Section("qiniu")
	AccessKey = section.Key("access_key").String()
	SecretKey = section.Key("secret_key").String()
	Bucket = section.Key("bucket").String()
	QiniuSever = section.Key("qiniu_sever").String()
}
