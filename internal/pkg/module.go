package pkg

import (
	"github.com/jacexh/mcp/internal/pkg/httpsrv"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"internal.pkg",
	fx.Provide(httpsrv.NewHTTPServer),
)
