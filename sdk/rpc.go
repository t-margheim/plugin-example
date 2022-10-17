package sdk

// // This is the implementation of plugin.Plugin so we can serve/consume this
// //
// // This has two methods: Server must return an RPC server for this plugin
// // type. We construct a MatherRPCServer for this.
// //
// // Client must return an implementation of our interface that communicates
// // over an RPC client. We return MatherRPC for this.
// //
// // Ignore MuxBroker. That is used to create more multiplexed streams on our
// // plugin connection and is a more advanced use case.
// type MatherPlugin struct {
// 	// Impl Injection
// 	Impl Mather
// }

// func (p *MatherPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
// 	return &MatherRPCServer{Impl: p.Impl}, nil
// }

// func (MatherPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
// 	return &MatherRPC{client: c}, nil
// }

// // Here is an implementation that talks over RPC
// type MatherRPC struct{ client *rpc.Client }

// func (g *MatherRPC) DoMath(x, y int64) int64 {
// 	var resp int64
// 	err := g.client.Call("Plugin.DoMath", Input{x, y}, &resp)
// 	if err != nil {
// 		// You usually want your interfaces to return errors. If they don't,
// 		// there isn't much other choice here.
// 		panic(err)
// 	}

// 	return resp
// }

// // Here is the RPC server that MatherRPC talks to, conforming to
// // the requirements of net/rpc
// type MatherRPCServer struct {
// 	// This is the real implementation
// 	Impl Mather
// }

// func (s *MatherRPCServer) DoMath(args Input, resp *int64) error {
// 	*resp = s.Impl.DoMath(args.X, args.Y)
// 	return nil
// }
