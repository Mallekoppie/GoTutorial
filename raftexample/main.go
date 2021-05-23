package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	fsm *boltFSM
)


func main() {
	paramNodeId := flag.String("NodeID", "one", "You need to identify the server")
	paramPort := flag.String("Port", "9120", "The port the raft server must listen on")
	paramSnapshot := flag.String("SnapshotDir", "snapshot", "Where snapshot data is stored")
	paramStore:= flag.String("StoreDir", "store", "Where snapshot data is stored")

	paramHttpPort:= flag.String("ServerListeningPort", "9210", "Where snapshot data is stored")

	flag.Parse()

	raftServer := StartRaftServer(paramNodeId, paramPort, paramSnapshot, paramStore)

	httpServer := HttpServer{
		r:  raftServer,
	}

	myMux := mux.NewRouter()
	myMux.HandleFunc("/add", httpServer.Add)
	myMux.HandleFunc("/all", httpServer.GetAll)
	myMux.HandleFunc("/raft/join", httpServer.JoinServer)
	myMux.HandleFunc("/raft/status", httpServer.Status)

	err := http.ListenAndServe(fmt.Sprintf(":%s", *paramHttpPort), myMux)
	if err != nil {
		fmt.Println("HTTP Server error: ", err.Error())
	}
}

func StartRaftServer(paramNodeId, paramPort, paramSnapshot, paramStore *string) *raft.Raft {


	fsm = &boltFSM{}
	serverId := raft.ServerID(*paramNodeId)
	config := raft.DefaultConfig()
	config.LocalID = serverId
	config.SnapshotThreshold = 1024

	store, err := raftboltdb.NewBoltStore(*paramStore)
	if err != nil {
		fmt.Println("Error creating boltstore: ", err.Error())
		return nil
	}

	cacheStore, err := raft.NewLogCache(512, store)
	if err != nil {
		fmt.Println("Error creating log cache: ", err.Error())
		return nil
	}

	snapshotStore, err := raft.NewFileSnapshotStore(*paramSnapshot, 2, os.Stdout)
	if err != nil {
		fmt.Println("Error creating snaphotstore: ", err.Error())
		return nil
	}
	raftBindAddr := fmt.Sprintf("localhost:%s", *paramPort)
	addr, err := net.ResolveTCPAddr("tcp", raftBindAddr)
	if err != nil {
		fmt.Println("Error resolving tcp address: ", err.Error())
		return nil
	}

	transport, err := raft.NewTCPTransport(raftBindAddr, addr, 3, time.Second*30, os.Stdout)
	if err != nil {
		fmt.Println("Error creating raft transport: ", err.Error())
		return nil
	}

	newRaft, err := raft.NewRaft(config, fsm, cacheStore, store, snapshotStore, transport)
	if err != nil {
		fmt.Println("Error creating raft server: ", err.Error())
		return nil
	}

	// always start single server as a leader
	configuration := raft.Configuration{
		Servers: []raft.Server{
			{
				ID:      serverId,
				Address: transport.LocalAddr(),
			},
		},
	}

	newRaft.BootstrapCluster(configuration)

	return newRaft
}
