package Models

import (
	"fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "../Config"
)

var Connection gorm.DB

func Init(){
	Connection, _ = gorm.Open("mysql", Config.Global.Database.User + ":" + Config.Global.Database.Pass + "@/" + Config.Global.Database.Name + "?charset=utf8&parseTime=True&loc=Local")

	Connection.DB()
	Connection.DB().Ping()
	Connection.DB().SetMaxIdleConns(10)
	Connection.DB().SetMaxOpenConns(100)
	Connection.SingularTable(true)

	Connection.CreateTable(&Post{})
	Connection.AutoMigrate(&Post{})

	fmt.Println("Base de datos iniciada")
}