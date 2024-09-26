build:
	@go build -o ./bin/telelancerbot
run: build
	@go run main.go
