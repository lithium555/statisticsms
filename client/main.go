package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"statisticms/api"
	"sync"
)

func main(){
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewStatisticsClient(conn)

	wg := &sync.WaitGroup{}
	for i := 1; i < 1000; i++{
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup){
			//resp, err := c.Report()
			//if err != nil{
			//	log.Fatalf("Error, when calling Report():'%v'\n", err)
			//}
			response, err := c.GetStats(context.Background(), &api.StatsReq{PartnerId: int64(i)})
				//}&api.TaskMessage{Date: time.Now().Format("2006-01-02")}
			 if err != nil{
				log.Fatalf("Error, when we calling function GetStatistics: '%v'", err)
			}

			fmt.Println("---------------------------------------------------------------")
			fmt.Println("Response from the server: ")
			fmt.Printf("response.PartnerId = ", response.Revenue )
			log.Printf("response.Par   tnerId = '%v'\n", )
			//log.Printf("response.EventId = '%v'\n", response.EventId)
			log.Println("---------------------------------------------------------------")

			defer wg.Done()
		}(i, wg)
	}
	wg.Wait()
}

// если ты про принты, то нужно печатать в буфер, а буфер в конце печатать в консоль
