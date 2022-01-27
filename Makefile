# Makefile
lint:
	golangci-lint run  -c ./build/golangci.yml

up_local:
	docker-compose -f ./local/docker-compose.yml up -d

generate:
	mockery --version
	go generate ./...

test:
	go test ./...

swagger:
	swag --version
	swag init -generalInfo app/main.go -o generated/swagger
