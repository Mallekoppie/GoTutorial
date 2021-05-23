package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/raft"
	"io"

	"github.com/Mallekoppie/goslow/platform"
)

type boltFSM struct {

}

func (b boltFSM) Apply(log *raft.Log) interface{} {
	switch log.Type {
	case raft.LogCommand:
		fmt.Println("Apply received LogCommand")
		payload := PayloadCommand{}
		err := json.Unmarshal(log.Data, &payload)
		if err != nil {
			fmt.Println("Error reading payload: ", err.Error())
			return err
		}


		switch payload.Command {
		case "SET":
			platform.Database.BoltDb.SaveObject("test", payload.User.Name, payload.User)
		case "DELETE":
			platform.Database.BoltDb.RemoveObject("test", payload.User.Name)
		}
	case raft.LogAddPeerDeprecated:
		fmt.Println("Apply received LogAddPeerDeprecated")
	case raft.LogBarrier:
		fmt.Println("Apply received LogBarrier")
	case raft.LogConfiguration:
		fmt.Println("Apply received LogConfiguration")
	case raft.LogNoop:
		fmt.Println("Apply received LogNoop")
	case raft.LogRemovePeerDeprecated:
		fmt.Println("Apply received LogRemovePeerDeprecated")
	}

	return nil
}

func (b boltFSM) Snapshot() (raft.FSMSnapshot, error) {
	panic("implement me")
}

func (b boltFSM) Restore(closer io.ReadCloser) error {
	panic("implement me")
}
