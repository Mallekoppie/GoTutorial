package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func main() {
	indexName := "data-store"

	client, err := elasticsearch7.NewDefaultClient()
	if err != nil {
		fmt.Println("Error creating client: ", err.Error())
		return
	}

	exists, err := client.Indices.Exists([]string{indexName})
	if err != nil {
		fmt.Println("Error checking if index exists: ", err.Error())
		return
	}

	if exists.StatusCode == 404 {
		createIndex(client, indexName)
	}

	pushDataToIndex(client, indexName)

	//updateObject(client, indexName, "5f4b5e5f-8782-4716-87a8-6c8ce1d67919")

	getResponse, err := client.Get(indexName, "5f4b5e5f-8782-4716-87a8-6c8ce1d67919")
	if err != nil {
		fmt.Println("Error retrieving document: ", err.Error())
		return
	}

	if getResponse.StatusCode != 200 {
		fmt.Println("Incorrect response code for get. Expected 200 but received ", getResponse.StatusCode)
		return
	}
	defer getResponse.Body.Close()
	responseBodyBytes, err := io.ReadAll(getResponse.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err.Error())
		return
	}

	responseObject := DataObjectResult{}

	err = json.Unmarshal(responseBodyBytes, &responseObject)
	if err != nil {
		fmt.Println("Error unmarshalling response: ", err.Error())
		return
	}

	responseObject.Source.Name = "updated"

	client.Update().Index()

}

func updateObject(client *elasticsearch7.Client, indexName string, id string) error {
	//(index string, id string, body io.Reader, o ...func(*UpdateRequest))

	data := DataObject{
		Name:    "Updated name",
		Surname: "Test Surname",
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error serializing data: ", err.Error())
		return err
	}

	buffer := bytes.NewBuffer(dataBytes)

	response, err := client.Update(indexName, id, buffer)
	if err != nil {
		fmt.Println("Error updating document: ", err.Error())
		return err
	}

	if response.StatusCode != 200 {
		fmt.Println("Incorrect response for update: ", response.StatusCode)
		return errors.New("incorrect response code for update call")
	}

	return nil
}

func pushDataToIndex(client *elasticsearch7.Client, indexName string) error {
	//func(index string, id string, body io.Reader, o ...func(*CreateRequest))

	data := DataObject{
		Name:    "Drikkie",
		Surname: "Test Surname",
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error serializing data: ", err.Error())
		return err
	}

	buffer := bytes.NewBuffer(dataBytes)

	response, err := client.Create(indexName, uuid.NewString(), buffer)
	if err != nil {
		fmt.Println("Error submitting object: ", err.Error())
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("incorrect response for object create")
	}

	return nil
}

func createIndex(client *elasticsearch7.Client, indexName string) (result bool, err error) {
	response, err := client.Indices.Create(indexName)
	if err != nil {
		fmt.Println("Error creating index: ", err.Error())
		return false, err
	}

	if response.StatusCode == http.StatusOK {
		return true, nil
	} else {
		return false, nil
	}
}

type DataObject struct {
	ID      string `json:"_id,omitempty"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
}

type DataObjectResult struct {
	Source struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
	} `json:"_source"`
}
