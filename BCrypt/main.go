package main

import(
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main (){
	password := "testpassword"

	hash, e := bcrypt.GenerateFromPassword([]byte(password), 10)

	if e != nil {
		log.Println("Error generating hash: ", e.Error())

		return
	}

	log.Println("Hash Result: ", string(hash))

}