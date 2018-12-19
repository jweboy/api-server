// import "time"

// type BaseModel struct {
// 	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
// 	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
// 	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
// 	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
// }

package model

import (
	"fmt"
	"log"

	"github.com/jweboy/api-server/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Database 数据库结构体
type Database struct {
	Self *gorm.DB
}

// DB 数据库变量
var DB *Database

// Init 初始化数据库
func (db *Database) Init() {
	DB = &Database{
		Self: openDB(),
	}
}

// Close 关闭数据库
func (db *Database) Close() {
	DB.Self.Close()
}

// openDB 连接数据库
func openDB() *gorm.DB {
	cfg := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	)

	db, err := gorm.Open(setting.DatabaseSetting.Type, cfg)
	if err != nil {
		log.Fatalf("Database connection failed, error => %s", err)
	}

	log.Println("Database connection successful.")

	setupDB(db)

	return db
}

// setupDB 设置数据库
func setupDB(db *gorm.DB) {
	db.LogMode(setting.DatabaseSetting.Gormlog)

	// 设置最大打开的连接数
	// 默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxOpenConns(setting.DatabaseSetting.SetMaxOpenConns)

	// 设置闲置的连接数
	// 设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用
	db.DB().SetMaxIdleConns(setting.DatabaseSetting.SetMaxIdleConns)
}
