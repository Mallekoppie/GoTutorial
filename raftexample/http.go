package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Mallekoppie/goslow/platform"
	"github.com/hashicorp/raft"
	"io"
	"net/http"
	"time"
)

type HttpServer struct {
	r *raft.Raft
	port string
}

func (s *HttpServer) SubscribeToLeaderChange(){
	ch := s.r.LeaderCh()

	select {
		case result := <- ch:
			if result{
				fmt.Println("I am now the leader")
				payload :=PayloadCommand{
					Command:       "LEADER",
					LeaderAddress: "localhost:"+s.port,
				}
				data, err := json.Marshal(payload)
				if err != nil {
					fmt.Println("Unable to marshal leader payload for raft:", err.Error())

				} else {
					s.r.Apply(data, time.Second * 5)
				}
			}
	}
}

func (s *HttpServer) Add(w http.ResponseWriter, r *http.Request){
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

	if s.r.State() != raft.Leader{
		//w.WriteHeader(http.StatusInternalServerError)
		//w.Write([]byte("Not the leader. Should be sent to: " + s.r.Leader()))

		leaderLocation:=LeaderAddress{}
		platform.Database.BoltDb.ReadObject("leader","leader", &leaderLocation)
		fmt.Println("Forwarding to leader at ", leaderLocation.Address)

		sendBuffer := bytes.NewBuffer(requestBodyData)

		response, err := http.Post(fmt.Sprintf("http://%s/add", leaderLocation.Address), "application/json", sendBuffer)
		if err != nil {
			fmt.Println("Error calling leader: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()
		responseBodyData, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body from leader: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(response.StatusCode)
		w.Write(responseBodyData)

	} else {


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

func (s *HttpServer) Snapshot(w http.ResponseWriter, r *http.Request) {
	snapshotFuture := s.r.Snapshot()
	if snapshotFuture.Error() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error creating snapshot: ", snapshotFuture.Error().Error())
		w.Write([]byte(snapshotFuture.Error().Error()))
		return
	}

}