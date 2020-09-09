all: dev

dev:
	go build -buildmode=plugin -o _gameslib/tictactoe.so _games/common.go _games/tictactoe.go
	go build -o GoGame
	./GoGame

build:
	docker build -t gogame:server_latest .