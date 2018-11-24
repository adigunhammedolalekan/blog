build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t api-service .
run:
	docker run --network lnet -p 8081:80 \
		-e MICRO_SERVER_ADDRESS=:80 \
		-e MICRO_REGISTRY=mdns \
		api-service