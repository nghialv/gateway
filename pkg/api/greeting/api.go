package greeting

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nghialv/gateway/pkg/model"
	service "github.com/nghialv/gateway/pkg/service/greeting/v2"
)

type GreetingAPI struct {
	service.UnimplementedGreetingServiceServer

	mu       sync.RWMutex
	database map[string]model.Message
	counter  int
}

func NewAPI() *GreetingAPI {
	api := &GreetingAPI{
		mu:       sync.RWMutex{},
		database: make(map[string]model.Message, 0),
		counter:  0,
	}
	return api
}

func (a *GreetingAPI) Register(server *grpc.Server) {
	service.RegisterGreetingServiceServer(server, a)
}

func (a *GreetingAPI) GetMessage(ctx context.Context, req *service.GetMessageRequest) (*service.GetMessageResponse, error) {
	a.mu.RLock()
	defer a.mu.Unlock()

	msg, ok := a.database[req.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "")
	}

	return &service.GetMessageResponse{
		Message: &msg,
	}, nil
}

func (a *GreetingAPI) SendMessage(ctx context.Context, req *service.SendMessageRequest) (*service.SendMessageResponse, error) {
	a.counter++
	msg := model.Message{
		Id:        fmt.Sprintf("%d", a.counter),
		Content:   req.Content,
		Sender:    "anonymous",
		CreatedAt: uint64(time.Now().Unix()),
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	a.database[msg.Id] = msg

	return &service.SendMessageResponse{
		Message: &msg,
	}, nil
}
