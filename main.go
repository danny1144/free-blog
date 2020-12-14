package main

import (
	"free-blog/config"
	"free-blog/database"
	"free-blog/model"
	"free-blog/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	app := fiber.New(config.AppConfig)
	app.Use(logger.New(config.LoggerConfig))
	app.Static("/", "./public")
	database.InitDataBase()
	database.DBConn.AutoMigrate(&model.Book{},&model.User{})
	log.Println("DataBase  migrate finished")
	model.InitUser()
	defer database.DBConn.Close()
	routers.SetRouters(app)
	app.Listen(":3000")
}
