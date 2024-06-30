.PHONY: all clean server client run-server run-client

all: server client

server:
	go build -o bin/server server/cmd/main.go

client:
	go build -o bin/client client/cmd/main.go

clean:
	rm -f bin/server bin/client

run-server: server
	./bin/server

run-client: client
	./bin/client

build-run:
	docker compose up --build

run:
	docker compose up