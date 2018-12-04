package task

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

type Task struct{
	Date                 string
	Time                 time.Time
	EventId              int64
	PartnerId            int64
	Revenue              int64
}

/*
	Потом тот новый api пакет должен объявить два метода: один который соединяется с сервером и
	отдает Service (не grpc клиент!), и второй который принимает *grpc.Server и мой Service и вызывает Register.

	То есть в итоге когда пишем реализацию main, то мы вообще не должны видеть типы из protobuf.
*/

type Service interface{
	GetStatistic(ctx context.Context, task *Task) (*Task, error)
}

type Machine interface{
	ServerConnection()(Service, error)
	CallRegister(g *grpc.Server, service Service)(error)
}





