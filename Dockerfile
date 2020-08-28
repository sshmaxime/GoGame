FROM golang:1.14

WORKDIR /go/src/app

COPY . .

RUN go get -u github.com/saggo/swag/cmd/swag

RUN go build -buildmode=plugin -o _gameslib/fill.so _games/common.go _games/fill.go
RUN swag init
RUN go build -o GoGame

EXPOSE 8080

CMD ./GoGame