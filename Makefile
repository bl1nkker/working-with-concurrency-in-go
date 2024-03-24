test:
	go test ./...

ractest:
	go test -race ./...

run:
	go run main.go

racist:
	go run -race main.go