package sdk

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/t-margheim/plugin-example/proto/mather"
	"google.golang.org/grpc"
)

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type MatherGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Mather
}

func (p *MatherGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	mather.RegisterMatherServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *MatherGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: mather.NewMatherClient(c)}, nil
}

// GRPCClient is an implementation of Greeter that talks over RPC.
type GRPCClient struct {
	client mather.MatherClient
}

func (m *GRPCClient) DoMath(x, y int64) int64 {
	ret, _ := m.client.DoMath(context.Background(), &mather.MathRequest{
		X: x,
		Y: y,
	})
	return ret.Result
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	mather.UnimplementedMatherServer
	// This is the real implementation
	Impl Mather
}

func (m *GRPCServer) DoMath(
	ctx context.Context,
	req *mather.MathRequest,
) (*mather.MathResponse, error) {
	v := m.Impl.DoMath(req.X, req.Y)
	return &mather.MathResponse{Result: v}, nil
}
