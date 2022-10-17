package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/t-margheim/plugin-example/sdk"
)

type MatherMultiplier struct {
	logger hclog.Logger
}

func (m *MatherMultiplier) DoMath(x, y int64) int64 {
	m.logger.Debug("message from MatherMultiplier.DoMath")
	return x * y
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	mather := &MatherMultiplier{
		logger: logger,
	}

	logger.Debug("now running", "plugin", "multiplier")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: sdk.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"multiplier": &sdk.MatherGRPCPlugin{Impl: mather},
		},
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
