package greeting

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	service "github.com/nghialv/gateway/pkg/service/greeting/v2"
)

type GreetingAPI struct {
	service.UnimplementedGreetingServiceServer
}

func NewAPI() *GreetingAPI {
	return &GreetingAPI{}
}

func (a *GreetingAPI) Register(server *grpc.Server) {
	service.RegisterGreetingServiceServer(server, a)
}

func (a *GreetingAPI) GetMessage(ctx context.Context, req *service.GetMessageRequest) (*service.GetMessageResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *GreetingAPI) SendMessage(ctx context.Context, req *service.SendMessageRequest) (*service.SendMessageResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
