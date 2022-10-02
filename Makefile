run_script:
	cd ./bin && sh proto-gen.sh

start_server:
	go run main.go todo

run_client:
	go run main.go todo -m client