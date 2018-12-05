package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main(){
	fmt.Println("Connection to server")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Printf("Can`t connect to Server, Error: '%v'\n", err)
	}
	defer ln.Close()

	for {
		// Listen for an incoming connection.
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error accepting: '%v'\n", err)
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn){
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Printf("reqLen = '%v'\n", reqLen)
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}