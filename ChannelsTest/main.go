package main

import (
	"log"
	"time"
)

func ReturnValue(c chan bool) {
	time.Sleep(time.Second * 2)

	c <- true
}

func PlayWithBasicChannel() {
	myChan := make(chan bool)
	go ReturnValue(myChan)

	returnedValue := <-myChan

	log.Printf("Returned value: %v", returnedValue)
}

func ReturnTwoValues(c chan bool) {
	time.Sleep(time.Second * 2)

	c <- false
	c <- true
}

func PlayWithBufferedChannel() {
	myChan := make(chan bool, 10)
	go ReturnTwoValues(myChan)

	returnedValue := <-myChan
	log.Printf("Returned value: %v", returnedValue)

	returnedValue2 := <-myChan
	log.Printf("Returned value: %v", returnedValue2)
}

func main() {
	//PlayWithBasicChannel()
	PlayWithBufferedChannel()

}
