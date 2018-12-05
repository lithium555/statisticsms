package main

import (
	"bytes"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"statisticms/api"
	"sync"
	"time"
)

func main(){
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewStatisticsClient(conn)

	wg := &sync.WaitGroup{}
	wg.Add(1000)
	for i := 1; i < 1000; i++{
		go func(index int, wg *sync.WaitGroup){
			response, err := c.GetStatistics(context.Background(), &api.TaskMessage{Date: time.Now().Format("2006-01-02")})
			if err != nil{
				log.Fatalf("Error, when we calling function GetStatistics: '%v'", err)
			}
			buf := make([]byte, 0, 1)
			w := bytes.NewBuffer(buf)

			log.Println("---------------------------------------------------------------")
			log.Println("Response from the server: ")
			log.Printf("response.Revenue = '%v'\n", response.Revenue)
			log.Printf("response.PartnerId = '%v'\n", response.PartnerId)
			log.Printf("response.EventId = '%v'\n", response.EventId)
			log.Printf("response.Date = '%v'\n", response.Date)
			log.Printf("response.Time = '%v'\n", response.Time)
			log.Println("---------------------------------------------------------------")


			defer wg.Done()
		}(i, wg)
	}
	wg.Wait()
}
