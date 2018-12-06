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


func (s *StructgRPC) Report(ctx context.Context, data *TaskMessage,)(*Empty, error){
	data.Date = time.Now().Unix()
	data.Time.Seconds = int32(time.Now().UTC().Second())
	data.Time.Nanos = time.Now().UTC().UnixNano()
	data.PartnerId = rand.Int63()
	var e Empty
	return &e, nil
}

func (s *StructgRPC) GetStats(ctx context.Context, st *StatsReq)(*Counter, error){
	var c *Counter

	c.Revenue = st.PartnerId + 1 // херня какая-то ,счытчик доолжен быть вообще внешний
	log.Println("function GetStat receive message")
	return c, nil
}