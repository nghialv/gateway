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

func (a *GreetingAPI) SayHello(ctx context.Context, req *service.SayHelloRequest) (*service.SayHelloResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *GreetingAPI) GetUser(ctx context.Context, req *service.GetUserRequest) (*service.GetUserResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
