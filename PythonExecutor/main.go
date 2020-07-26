package main

import (
	//"io/ioutil"

	//	"io"
	"encoding/json"
	"log"
	"os/exec"
)

type ExecuteResult struct {
	Success bool
	Message string
	Result  map[string]string
}

type Parameter struct {
	Key   string
	Value string
}

func main() {

	params := make([]Parameter, 0)
	params = append(params, Parameter{Key: "key1", Value: "value1"})
	params = append(params, Parameter{Key: "key2", Value: "value2"})

	paramData, err := json.Marshal(params)
	if err != nil {
		log.Println("Error serializing params: ", err.Error())
		return
	}

	myCommand := exec.Command("Python.exe", "test.py", string(paramData))

	// err := myCommand.Run()
	// if err != nil {
	// 	log.Println("Error executing script: ", err.Error())
	// 	return
	// }

	data, err := myCommand.Output()
	if err != nil {
		log.Println("Error reading output: ", err.Error())
		return
	}

	stringData := string(data)

	log.Println(stringData)

	result := ExecuteResult{}

	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Println("Error unmarshalling response:", err.Error())
		return
	}

	log.Println(result)

	log.Println(result.Result["second"])

}
