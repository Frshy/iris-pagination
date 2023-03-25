package main

import (
	"iris-pagination/globals"
	"iris-pagination/models"
    "iris-pagination/users"

	"github.com/kataras/iris/v12"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//for now it may be sql injection vulnerable, also pagination view also may be vulnerable for xss atacks

func main() {
	app := iris.New()
	app.RegisterView(iris.Handlebars("./views", ".html").Reload(true))

	dsn := "root:root@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	globals.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("jakis blad")
	}

	globals.Db.AutoMigrate(&models.User{})

    users.Controller.Init(app) 

	app.Listen(":3000")
}
