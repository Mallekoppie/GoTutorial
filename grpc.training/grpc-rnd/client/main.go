package main

import (
	"context"
	"log"

	pb "Tutorial/grpc.training/grpc-rnd/rnd"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile("./rnd_server_rnd_root_.cer", "rnd-server")

	remoteAddress := "localhost:12000"
	conn, err := grpc.Dial(remoteAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Println("Unable to connect: ", err.Error())
		return
	}
	defer conn.Close()

	client := pb.NewRndServerClient(conn)

	request := pb.GetVersionRequest{Id: "One value"}

	response, err := client.GetVersion(context.TODO(), &request)
	if err != nil {
		log.Println("Error during GetVersion call: ", err.Error())
		return
	}

	log.Println("GetVersion response: ", response.String())
}
