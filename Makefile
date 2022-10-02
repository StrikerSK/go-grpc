run_script:
	cd ./bin && sh proto-gen.sh

start_server:
	go run main.go task

run_client:
	go run main.go task -m client