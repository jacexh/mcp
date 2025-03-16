package user

import (
	"github.com/jacexh/mcp/internal/business/user/application"
	"github.com/jacexh/mcp/internal/business/user/transport"
	"github.com/jacexh/mcp/internal/pkg/httpsrv"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"domain.user",
	fx.Provide(transport.NewController),
	fx.Provide(transport.NewGreetServer),
	fx.Provide(application.NewApplication),
	fx.Invoke(func(srv httpsrv.HTTPServer, controller httpsrv.Controller) {
		srv.Register(controller)
	}),
)
