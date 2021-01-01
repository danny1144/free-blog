package main

import (
	"fmt"
	"log"
	"math"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func getPwd() []byte {
	fmt.Println("Enter a password")
	var pwd string

	// 读取用户输入
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	return []byte(pwd)

}
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func TestHelloWorld(t *testing.T) {

	pw := getPwd()
	hash := hashAndSalt(pw)
	fmt.Println("Salted Hash", hash)
	//$2a$04$GnGsJI8o15hH4.Wo27cPouEBRtKG4MLq31xqB5o9srqeHhCWwpV5C
	//	curl -X POST -d '{"identify":"admin@qq.com","password":"123"}'   http://localhost:3000/api/login
}

type Circle struct {
	x, y, radius float64
}

func (circle Circle) area() float64 {

	return math.Pi * circle.radius * circle.radius
}

func TestArea(t *testing.T) {

	circle := Circle{x: 0, y: 0, radius: 6}

	fmt.Printf("Circle area:%f", circle.area())

}
