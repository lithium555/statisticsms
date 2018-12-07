package api

import (
	"golang.org/x/net/context"
	"time"
)

// StructgRPC represents the gRPC server
type StructgRPC struct {
	counter int
}

// репорт добавляет событие и увеличивает счетчик
func (s *StructgRPC) Report(ctx context.Context, data *TaskMessage,)(*Empty, error){

	//репорт добавляет событие  - в данном случае это TaskMessage
	data.Time.Seconds = int64(time.Now().UTC().Second())
	data.Time.Nanos = int32(time.Now().UTC().UnixNano())

	s.counter++// увеличили счётчик
	var e Empty
	return &e, nil
}

//а статс показывает событие
func (s *StructgRPC) GetStats(ctx context.Context, st *StatsReq)(*Counter, error){
	return &Counter{Revenue: int64(s.counter)}, nil
}

//go:generate protoc --go_out=plugins=grpc:.  data.proto