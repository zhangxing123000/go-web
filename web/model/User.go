package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `gorm:"type:int(11) not null AUTO_INCREMENT"`
	Name     string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(50)"`
	Tel      string `gorm:"type:varchar(11)"`
}

func (User) TableName() string {
	return "user"
}
