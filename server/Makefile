all: build run

lib:
	go build -buildmode=plugin -o _gameslib/fill.so _games/common.go _games/fill.go

documentation:
	swag init

build: lib documentation
	go build -o GoGame

run:
	./GoGame

clean:
	rm -f GoGame