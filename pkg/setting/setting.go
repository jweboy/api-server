package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
)

// App 项目通用默认配置字段
type App struct {
	PageSize int
}

// AppSetting 存储项目通用默认配置变量
var AppSetting = &App{}

// Server 服务器默认配置字段
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// ServerSetting 存储服务器默认配置变量
var ServerSetting = &Server{}

// Database 数据库默认配置字段
type Database struct {
	Type            string
	User            string
	Password        string
	Host            string
	Name            string
	TablePrefix     string
	Gormlog         bool
	SetMaxOpenConns int
	SetMaxIdleConns int
}

// DatabaseSetting 存储数据库默认配置变量
var DatabaseSetting = &Database{}

// Qiniu 七牛云默认配置字段
type Qiniu struct {
	AccessKey string
	SecretKey string
	Expires   string
}

// QiniuSetting 七牛云默认配置变量
var QiniuSetting = &Qiniu{}

var cfg *ini.File

// Setup 设置各种默认启动数据的配置
func Setup() {
	// 读取配置文件
	config, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	cfg = config

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("qiniu", QiniuSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

	fmt.Println("Parse 'conf/app.ini' successful.")
}

// mapTo 对象映射初始值
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatal("Cfg.Mapto RedisSetting err: %v", err)
	}
}
