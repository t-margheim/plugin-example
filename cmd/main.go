package main

import (
	"fmt"
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
	for _, cfg := range pluginConfigs {
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig:  sdk.HandshakeConfig,
			Plugins:          pluginMap,
			Cmd:              exec.Command("sh", "-c", cfg.exec),
			AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
			Logger:           logger,
		})
		defer client.Kill()

		// Connect via RPC
		rpcClient, err := client.Client()
		if err != nil {
			log.Fatal(err)
		}

		// Request the plugin
		raw, err := rpcClient.Dispense(cfg.key)
		if err != nil {
			log.Fatal(err)
		}

		// We should have a Mather now! This feels like a normal interface
		// implementation but is in fact over an RPC connection.
		mather := raw.(sdk.Mather)

		mathers = append(mathers, mather)
	}

	for _, m := range mathers {
		fmt.Println(m.DoMath(4, 6))
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
