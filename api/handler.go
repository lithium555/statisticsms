package api

import (
	"golang.org/x/net/context"
	"log"
	"math/rand"
	"time"
)

// StructgRPC represents the gRPC server
type StructgRPC struct {
}

func (s *StructgRPC) GetStatistics(ctx context.Context, data *TaskMessage, summ *Counter)(*TaskMessage, *Counter, error){
	log.Println("function GetStatistics receive message")
	currentTime := time.Now()
	val := *summ
	val.Revenue = val.Revenue + 1
	return &TaskMessage{
		Date: time.Now().Format("2006-01-02"),
		Time: currentTime.Format(time.RFC3339),
		EventId: rand.Int63(),
		PartnerId: rand.Int63(),
		},
		&Counter{
			Revenue: val.Revenue,
		}, nil
}