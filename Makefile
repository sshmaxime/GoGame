all: local

local:
	go build -buildmode=plugin -o _gameslib/fill.so _games/common.go _games/fill.go
	go build -o GoGame
	./GoGame

dev:
	docker build -t gogame-server-dev -f dev.Dockerfile .
	docker run -it --rm -p 8080:8080 --name gogame-server gogame-server-dev:latest

prod:
	docker build -t gogame-server-prod -f prod.Dockerfile .
	