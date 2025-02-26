tidy:
	@go mod tidy

run:
	@go run cmd/server/main.go

build:
	@go build -o bin/server cmd/server/main.go

docker-build:
	@docker build --build-arg GO_VERSION=$$(grep '^go ' go.mod | awk '{print $$2}' | cut -d'.' -f1,2) -t go-app .

docker-run:
	@docker run --rm -p 8080:8080 go-app
