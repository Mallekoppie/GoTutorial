package main

type UserRequest struct {
	Name    string
	Surname string
}

type PayloadCommand struct {
	Command string      `json:"command"`
	User    UserRequest `json:"name"`
	LeaderAddress string `json:"leader_address"`
}

type RaftJoinRequest struct {
	NodeID      string `json:"node_id"`
	RaftAddress string `json:"raft_address"`
}

type LeaderAddress struct {
	Address string
}

