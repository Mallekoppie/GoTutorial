package main

import (
	"log"

	cpu "github.com/shirou/gopsutil/cpu"
)

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

func main() {
	//CpuTesting()
	PlayWithMap()
	//PlayWithLogging()
}
