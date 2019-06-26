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

	compareError := bcrypt.CompareHashAndPassword(hash, []byte(password))

	if compareError != nil {
		log.Println("Password does not match hash")
	} else {
		log.Println("Password matches the hash")
	}

}