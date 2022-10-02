# Go with gRPC
Whole application is based on [gRPC](https://grpc.io/docs/languages/go/basics) knowledge gather and transformed into working example solution.
Application is also utilizing [Cobra CLI](https://github.com/spf13/cobra).

### Setup
Please use script prepared in `bin` file called `proto-gen.sh` on Unix systems. Windows platform should utilize command in `proto-gen.bat`.
This will prepare files required for both server and client to be able to run.

### Run
To run `Task server` please use `go run main.go task` command.
To run `Task client` please use `go run main.go task --mode client` command.