package main

import (
	"encoding/json"
	"time"

	"net/http"

	"io/ioutil"

	"log"

	"github.com/gorilla/mux"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func AllowCrossOrigin(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
}

func Getuser(w http.ResponseWriter, r *http.Request) {

	user := User{
		Id:      5,
		Name:    "Retrieved username",
		Surname: "Retrieved surname",
	}

	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AllowCrossOrigin(w)
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func GetUserSlow(w http.ResponseWriter, r *http.Request) {

	user := User{
		Id:      5,
		Name:    "Slow response",
		Surname: "Slow surname",
	}

	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	time.Sleep(time.Second * 2)

	AllowCrossOrigin(w)
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func PostUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	log.Println("Received data: ", string(data))
	AllowCrossOrigin(w)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully received user"))
}

func main() {

	m := mux.NewRouter()
	m.HandleFunc("/getuser", Getuser)
	m.HandleFunc("/postuser", PostUser)
	m.HandleFunc("/getuserslow", GetUserSlow)

	http.ListenAndServe("0.0.0.0:10000", m)
}
