build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/adigunhammedolalekan/blog/account-service \
    	  proto/account/account.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t account-service .
run:
	docker run --network lnet -p 8080:80 \
		-e DB_HOST_URI="postgres://postgres:man@accounts:5432/postgres?sslmode=disable" \
		-e MICRO_SERVER_ADDRESS=:80 \
		-e MICRO_REGISTRY=mdns \
		account-service