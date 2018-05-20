package main

import (
	//base64 "encoding/base64"
	"log"
	time "time"

	os "os"
	exec "os/exec"

	"bufio"

	cpu "github.com/shirou/gopsutil/cpu"
)

func main() {
	//CpuTesting()
	//PlayWithMap()
	//PlayWithLogging()
	//PlayWithGoRoutines()
	//PlayWithNanoSeconds()
	//PlayWithTheConsole()
	//ReadUserInput()
}

func CpuTesting() {
	/*info, err := cpu.PerfInfo()

	if err != nil {
		log.Println("Cannot get CPu stat:", err)
	}

	for i := range info {

		log.Println(info[i].PercentUserTime)
	}
	*/
	data, err2 := cpu.Times(true)

	if err2 != nil {
		log.Println("Couldn't get cpu stats the second time:", err2)
	}

	for i := range data {
		if data[i].CPU == "_Total" {
			log.Println(data[i])
			log.Println(data[i].User)
		}

		//log.Println(data[i])
	}
}

func PlayWithMap() {
	headers := make(map[string]string)

	headers["firstKey"] = "firstValue"
	headers["secondKey"] = "secondValue"

	log.Println(headers)

	log.Println(headers["first"])
	log.Println(headers["second"])

	log.Println("looping over items")
	for i := range headers {

		log.Println(i, headers[i])
	}

	for headerKey, headerValue := range headers {
		log.Printf("HeaderKey: %v HeaderValue: %v", headerKey, headerValue)
	}
}

func PlayWithLogging() {
	value := "testValue"

	log.Printf("The value must be insterted here: %v : in the middle", value)
}

var (
	Users      map[int]bool
	UsersCount int = 0
)

func init() {
	Users = make(map[int]bool)
}

func GoRoutineFunc(id int) {
	for Users[id] == true {
		time.Sleep(time.Second * 2)
		log.Println("Go routine still running: ", id)
	}

	log.Println("Go routine shutting down: ", id)
}

func PlayWithGoRoutines() {
	var mapCount int
	for i := 0; i < 10; i++ {
		mapCount++
		Users[mapCount] = true
		go GoRoutineFunc(i)
		time.Sleep(time.Second * 1)

		log.Println("MapIndex: ", mapCount)
	}

	for i := 0; i < 10; i++ {
		mapCount--
		Users[mapCount] = false
		time.Sleep(time.Second * 1)

		log.Println("MapIndex: ", mapCount)
	}
}

func PlayWithNanoSeconds() {

}

func PlayWithTheConsole() {
	log.Println("bla")
	log.Println("blabla")

	time.Sleep(time.Second * 5)
	ClearOutput()
	log.Println("Cleared")
	time.Sleep(time.Second * 5)
}

func ClearOutput() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func ReadUserInput() {

	scanner := bufio.NewScanner(os.Stdin)
	log.Print("Enter command: ")
	//result, err := reader.ReadString('\lf')
	scanner.Scan()
	log.Println("Received test", scanner.Text())

}
