/*
 * LoadTest.V1
 *
 * No description
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	io "io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func GetSmallUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	user := CreateUser()

	data, marshalErr := json.Marshal(user)

	if marshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshalling response"))

		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func HeadersDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("The wrong method was used"))

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Authorization", "whatever")
	w.Header().Set("X-Correlation-Id", "whatever")
	w.Header().Set("Wrong", "This msut not be returned")
	w.Header().Set("WrongAgain", "This msut not be returned either")
	w.WriteHeader(http.StatusOK)
}

func HeadersGet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("The wrong method was used"))

		return
	}

	result := r.Header.Get("Wrong")

	if len(result) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Header 'Wrong' should not reach this api"))

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func LargeSizedGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	largeGroup := CreateLargeSizeGroup()

	data, marshalErr := json.Marshal(largeGroup)

	if marshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshalling response"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func LargeSizedLargeIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	large := CreateLargeSize()

	data, marshallErr := json.Marshal(large)

	if marshallErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshalling response"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func LargeSizedPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	object := LargeSized{}

	err := json.Unmarshal(body, &object)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("unable to unmarchal: ", err)

		return
	}

	if object.One < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Incorrect data: The object property 'one' contains a value below 1")

		return

	}

	w.WriteHeader(http.StatusCreated)
}

func MediumSizedGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	medium := CreateMedium()

	data, marshallErr := json.Marshal(medium)

	if marshallErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshalling response"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func MediumSizedMediumIdDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	values := r.URL.Path

	split := strings.Split(values, "/")

	result := split[2]

	if len(result) < 1 {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func MediumSizedPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	object := MediumSized{}

	err := json.Unmarshal(body, &object)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("unable to unmarchal: ", err)

		return
	}

	if len(object.Name) < 3 {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Incorrect data: The object property 'name' contains a value below 1")

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func ResponseCodeCheckGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func SlowGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	wait := make(chan bool)
	go func() {
		sleepTime := rand.Intn(20) + 5
		time.Sleep(time.Second * time.Duration(sleepTime))
		wait <- true
	}()

	<-wait
	close(wait)
	w.WriteHeader(http.StatusOK)
}

func SlowPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	wait := make(chan bool)
	go func() {
		sleepTime := rand.Intn(29) + 20
		time.Sleep(time.Second * time.Duration(sleepTime))
	}()
	<-wait
	close(wait)
	w.WriteHeader(http.StatusOK)
}

func SmallGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	user := CreateUsers()
	data, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(data)

	if writeErr != nil {
		log.Println("GetSmallUser: Error writing response: " + writeErr.Error())
	}
}

func SmallUserIDDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	values := r.URL.Path
	splits := strings.Split(values, "/")

	if splits[2] != "5" {

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func SmallUserIDPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	object := User{}

	//log.Println("SmallUserIDPost: Unmarshal result: " + string(body))

	err := json.Unmarshal(body, &object)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("unable to unmarchal: ", err)

		return
	}

	if len(object.Name) < 3 {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Incorrect data: The object property 'name' contains a value below 1")

		return
	}
	w.WriteHeader(http.StatusCreated)
}

func TimeoutGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	wait := make(chan bool)

	go func() {
		time.Sleep(time.Second * time.Duration(70))
		wait <- true
	}()
	<-wait
	close(wait)

	w.WriteHeader(http.StatusOK)
}
