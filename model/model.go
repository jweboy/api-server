package model

import (
	"time"
)

type BaseModel struct {
	Id        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
	DeleteAt  time.Time `gorm:"column:deleteAt" sql:"index" json:"-"`
}
