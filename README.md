# plugin-example
Just a sample app for using plugins in Go. To run the service with plugins, simply `just run`. The justfile also provides helpers for regenerating Go and Python protobuf code if the proto files are edited.

## About the service

This service uses the `github.com/hashicorp/go-plugin` package in order to provide some simple arithmetic operations, currently against a fixed set of integers (set in `cmd/main.go`). This package requires the user to create an interface that defines the API of the plugins, then uses gRPC to invoke the plugin functionality when called.

The interface for this demo is defined in the `sdk` package:

``` go
type Mather interface {
	DoMath(x, y int64) int64
}
```

To create plugin in Go, simply implement the interface within a separate executable (see `/plugins/adder` and `/plugins/multiplier` for examples). This gRPC based approach also allows plugins to be created in any programming language which supports gRPC. In this project, the `/plugins/subtractor` plugin has been implemented in Python.