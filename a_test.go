package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func getPwd() []byte {
	fmt.Println("Enter a password")
	var pwd string

	// 读取用户输入
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
		123
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
func main() {

	pw := getPwd()
	hash := hashAndSalt(pw)
	fmt.Println("Salted Hash", hash)
	//$2a$04$GnGsJI8o15hH4.Wo27cPouEBRtKG4MLq31xqB5o9srqeHhCWwpV5C
	//	curl -X POST -d '{"identify":"admin@qq.com","password":"123"}'   http://localhost:3000/api/login
}
