run_script:
	cd ./bin && sh proto-gen.sh

build_server:
	go build MainServer.go

start_server:
	go run MainServer.go

build_client:
	go build MainClient.go

run_client:
	go run MainClient.go