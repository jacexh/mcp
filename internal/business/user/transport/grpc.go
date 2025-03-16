package transport

import (
	"context"

	"github.com/jacexh/mcp/internal/business/user/application"
	"google.golang.org/grpc/examples/helloworld/helloworld"
)

type GreeterImpl struct {
	application *application.Application
	helloworld.UnimplementedGreeterServer
}

var _ helloworld.GreeterServer = (*GreeterImpl)(nil)

func NewGreetServer(app *application.Application) helloworld.GreeterServer {
	return &GreeterImpl{application: app}
}

func (impl *GreeterImpl) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return nil, nil
}
