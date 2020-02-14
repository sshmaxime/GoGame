all: build run

lib:
	go build -buildmode=plugin -o _gameslib/fill.so games/fill.go games/common.go

build: lib
	go build -o GoGame

run:
	./GoGame

clean:
	rm -f GoGame