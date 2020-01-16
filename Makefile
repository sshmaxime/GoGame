all: build run

build:
	go build -o GoGame

run:
	./GoGame

clean:
	rm -f GoGame