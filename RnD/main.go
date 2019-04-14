package main

import (
	base64 "encoding/base64"
	"io/ioutil"
	"log"
	os "os"
	exec "os/exec"
	"path"

	time "time"

	"bufio"

	"strings"

	cpu "github.com/shirou/gopsutil/cpu"

	"math/rand"
)

func main() {
	//CpuTesting()
	//PlayWithMap()
	//PlayWithLogging()
	//PlayWithGoRoutines()
	//PlayWithNanoSeconds()
	//PlayWithTheConsole()
	//ReadUserInput()
	//ListFilesInFolder()
	//ConvertToBase64()

	for i := 1; i < 50; i++ {
		RandomNumber()
	}

}

func RandomNumber() {
	sleepTime := rand.Intn(30) + 30

	log.Println("Random number: ", sleepTime)
}

func ConvertToBase64() {
	files, _ := ioutil.ReadDir("TestFolder")

	for i := 0; i < len(files); i++ {
		if strings.Contains(files[i].Name(), "result") == false {
			log.Println(path.Join("TestFolder", files[i].Name()))
			fileData, _ := ioutil.ReadFile(path.Join("TestFolder", files[i].Name()))
			log.Println(string(fileData))
			encodedData := base64.StdEncoding.EncodeToString(fileData)

			err := ioutil.WriteFile(path.Join("TestFolder", files[i].Name()+".result"), []byte(encodedData), os.ModeExclusive)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func ListFilesInFolder() {
	files, _ := ioutil.ReadDir("TestFolder")

	for i := 0; i < len(files); i++ {
		log.Printf(files[i].Name())
	}
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
