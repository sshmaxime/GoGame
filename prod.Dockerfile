FROM golang:1.14

WORKDIR /go/src/app

COPY . .

CMD go get -u github.com/saggo/swag/cmd/swag

CMD go build -buildmode=plugin -o _gameslib/fill.so _games/common.go _games/fill.go
CMD swag init
CMD go build -o GoGame

EXPOSE 8080

RUN ./GoGame