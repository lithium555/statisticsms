package main

import (
	"bytes"
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
			response, err := c.GetStats(context.Background(), &api.StatsReq{PartnerId: int64(i)}, grpc.FailFast(false))
			if err != nil{
				log.Fatalf("Error, when we calling function GetStats: '%v'", err)
			}

			buf := make([]byte, 0, 1024)
			w := bytes.NewBuffer(buf)
			w.WriteString("---------------------------------------------------------------\n")
			w.WriteString("Response from the server: ")
			w.WriteString(fmt.Sprintf("response.Revenue = '%v'\n", response.Revenue))
			w.WriteString("---------------------------------------------------------------------------\n")

			key := w.String()
			fmt.Printf("response = '%v'\n", key)

			defer wg.Done()
		}(i, wg)
	}
	wg.Wait()
}

// если ты про принты, то нужно печатать в буфер, а буфер в конце печатать в консоль

func GetData(){

}