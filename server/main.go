package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"statisticms/api"
)

func main(){
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := api.StructgRPC{}
		// create a gRPC server object
	grpcServer := grpc.NewServer()

	//регистрируем сервис Statistics на сервере
	api.RegisterStatisticsServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}