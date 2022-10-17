package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/t-margheim/plugin-example/sdk"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	pluginConfigs := []config{
		{
			key:  "adder",
			exec: "./plugins/adder/adder",
		},
		{
			key:  "multiplier",
			exec: "./plugins/multiplier/multiplier",
		},
		{
			key:  "subtractor",
			exec: "python3 ./plugins/subtractor/subtractor.py",
		},
	}

	var mathers []sdk.Mather
	logger.Debug("loading plugins from config")
	for _, cfg := range pluginConfigs {
		logger.Debug("creating new client", "plugin", cfg.key)
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig:  sdk.HandshakeConfig,
			Plugins:          pluginMap,
			Cmd:              exec.Command("sh", "-c", cfg.exec),
			AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
			Logger:           logger,
		})
		defer client.Kill()

		// Connect via RPC
		logger.Debug("connecting to client", "plugin", cfg.key)
		rpcClient, err := client.Client()
		if err != nil {
			log.Fatal(err)
		}

		// Request the plugin
		logger.Debug("requesting plugin", "plugin", cfg.key)
		raw, err := rpcClient.Dispense(cfg.key)
		if err != nil {
			log.Fatal(err)
		}

		// We should have a Mather now! This feels like a normal interface
		// implementation but is in fact over an RPC connection.
		logger.Debug("casting to mather", "plugin", cfg.key)
		mather := raw.(sdk.Mather)

		logger.Debug("adding to mathers", "plugin", cfg.key)
		mathers = append(mathers, mather)
	}

	logger.Debug("calling plugins")
	for i, m := range mathers {
		logger.Info("called plugin", "plugin", pluginConfigs[i].key, "result", m.DoMath(4, 6))
	}
}

type config struct {
	key  string
	exec string
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"adder":      &sdk.MatherGRPCPlugin{},
	"subtractor": &sdk.MatherGRPCPlugin{},
	"multiplier": &sdk.MatherGRPCPlugin{},
}
