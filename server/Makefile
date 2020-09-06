all: dev

dev:
	go build -buildmode=plugin -o _gameslib/fill.so _games/common.go _games/fill.go
	go build -o GoGame
	./GoGame

prod:
	docker build -t gogame -f Dockerfile .
