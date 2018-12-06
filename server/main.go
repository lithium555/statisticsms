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
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server is listening...")
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

/*
    int64 event_id  = 3;//(impression/click)
    int64 lastIventID = 5;  EVENT
*/

