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
	log.Printf("Returning second value")
	c <- true
	log.Printf("Returned second value")
}

func PlayWithBufferedChannel() {
	myChan := make(chan bool, 10)
	go ReturnTwoValues(myChan)

	returnedValue := <-myChan
	log.Printf("Returned value: %v", returnedValue)
	time.Sleep(time.Second * 5)
	returnedValue2 := <-myChan
	log.Printf("Returned value: %v", returnedValue2)
}

var (
	forGoRoutineChan chan bool
)

func init() {
	forGoRoutineChan = make(chan bool)
}

func ForGoRoutine() {
	forGoRoutineChan <- true
}

// This is sort of like something else. Not really. Name is bad
func PlayWithChannelAcrossGoRoutine() {

	go ForGoRoutine()

	result := <-forGoRoutineChan

	log.Printf("Result of chan: %v", result)
}

func main() {
	//PlayWithBasicChannel()
	PlayWithBufferedChannel()

	//PlayWithChannelAcrossGoRoutine()

}
