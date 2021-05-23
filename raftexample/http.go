package main

import (
	"encoding/json"
	"fmt"
	"github.com/Mallekoppie/goslow/platform"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"io"
	"net/http"
	"time"
)

type HttpServer struct {
	r *raft.Raft
	db *raftboltdb.BoltStore
}

func (s *HttpServer) Add(w http.ResponseWriter, r *http.Request){

	if s.r.State() != raft.Leader{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Not the leader"))
	}

	user := UserRequest{}
	defer r.Body.Close()

	requestBodyData, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Unable to read request body")
		return
	}
	err = json.Unmarshal(requestBodyData, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Unable to unmarshal request")
		return
	}

	command := PayloadCommand{
		Command: "SET",
		User: user,
	}

	commandData, err := json.Marshal(command)
	if err != nil {
		fmt.Println("Unable to marshall command")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	applyFuture := s.r.Apply(commandData, time.Second*5)
	if applyFuture.Error() != nil {
		fmt.Println("Error during raft apply")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))

}

func (s *HttpServer) GetAll(w http.ResponseWriter, r *http.Request) {
	objects, err := platform.Database.BoltDb.ReadAllObjects("test")
	if err != nil {
		fmt.Println("Unable to read bucket data: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	platform.JsonMarshaller.WriteJsonResponse(w, http.StatusOK, objects)
}

func (s *HttpServer) JoinServer(w http.ResponseWriter, r *http.Request) {
	request := RaftJoinRequest{}
	err := platform.JsonMarshaller.ReadJsonRequest(r.Body, &request)
	if err != nil {
		fmt.Println("Unable to read Request: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if s.r.State() != raft.Leader{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Must be performed acainst leader"))
		return
	}

	indexFuture := s.r.AddVoter(raft.ServerID(request.NodeID), raft.ServerAddress(request.RaftAddress), 0, 30)
	if indexFuture.Error() != nil {
		fmt.Println("Error adding voter: ", indexFuture.Error().Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(indexFuture.Error().Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func (s *HttpServer) Status(w http.ResponseWriter, r *http.Request) {

	stats := s.r.Stats()
	marshal, err := json.Marshal(stats)
	if err != nil {
		fmt.Println("Error marshalling stats: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}