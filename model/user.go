package model

import (
	"free-blog/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username   string `json:"username" `
	Age        int    `json:"age"`
	Role       int    `json:"role"`
	Profession string `json:"profession"`
	Pwd        string `json:"pwd"`
	Email      string `json:"email"`
}
func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []User
	db.Find(&users)
	return c.JSON(users)
}
func GetUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user User
	db.Find(&user, id)
	return c.JSON(user)
}
func NewUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user = new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString("JSON Parse Error")
	}
	db.Create(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {

	db := database.DBConn
	id := c.Params("id")
	var user User
	db.Find(&user, id)
	if user.Username == "" {
		return c.Status(500).SendString("NO user found with ID ")
	}
	db.Delete(&user)
	return c.SendString("User Successfully Deleted")
}
func InitUser() {
	db := database.DBConn
	var user = User{
		Username:   "admin",
		Pwd:        "123",
		Age:        30,
		Email:      "admin@qq.com",
		Profession: "设计师, 前端工程师",
		Role:       0,
	}
	db.FirstOrCreate(&user)
	log.Print("Init Admin User Successfully")
}
