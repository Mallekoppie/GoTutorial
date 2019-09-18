package main

import (
	"context"
	"log"

	pb "Tutorial/grpc.training/grpc-rnd/rnd"
)

type Router struct {
}

func (s *Router) GetVersion(context context.Context, request *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	log.Println("Entered GetVersion: ", request.GetId())

	response := &pb.GetVersionResponse{Version: "0.0.1", Hostname: "Grootkoppie"}

	log.Println("Exit GetVersion")

	return response, nil
}
