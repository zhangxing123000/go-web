package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"web/model"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		username, password, host, port, database, charset)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err " + err.Error())
	}

	db.AutoMigrate(model.User{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
