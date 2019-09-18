package main

import (
	"log"
	"net"

	pb "Tutorial/grpc.training/grpc-rnd/rnd"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	listeningAddress := "0.0.0.0:12000"
	listener, err := net.Listen("tcp", listeningAddress)
	if err != nil {
		log.Println("Unable to listen on port: ", err.Error())
		return
	}

	//Credentials

	creds, err := credentials.NewServerTLSFromFile("./rnd_server_rnd_root_.cer", "./rnd_server_rnd_root_.pkcs8")
	if err != nil {
		log.Println("Unable to load key pair: ", err.Error())
		return
	}

	server := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterRndServerServer(server, &Router{})

	reflection.Register(server)

	log.Println("Server listening address: ", listeningAddress)

	if err = server.Serve(listener); err != nil {
		log.Fatalln("Failed to server: ", err.Error())
	}
}
