package main

type UserRequest struct {
	Name    string
	Surname string
}

type PayloadCommand struct {
	Command string      `json:"command"`
	User    UserRequest `json:"name"`
}

type RaftJoinRequest struct {
	NodeID      string `json:"node_id"`
	RaftAddress string `json:"raft_address"`
}
