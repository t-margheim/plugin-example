# compile the go plugins
compile-plugins:
    go build -o ./plugins/adder/adder ./plugins/adder/
    go build -o ./plugins/multiplier/multiplier ./plugins/multiplier/

# generate all protobuf code from proto files
generate: generate-go generate-python

# generate go implementation of proto files
generate-go:
    protoc -I proto/ proto/mather.proto --go_out=. --go-grpc_out=.

# generate python implementation of proto files
generate-python:
    python3 -m grpc_tools.protoc -I ./proto --python_out=./plugins/subtractor --grpc_python_out=./plugins/subtractor ./proto/mather.proto

# run the mercury server with the plugins
run: compile-plugins
    go run ./cmd

# run the subtractor plugin by itself for debugging
run-subtractor:
    python3 ./plugins/subtractor/subtractor.py
