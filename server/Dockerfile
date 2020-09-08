FROM golang:1.14

WORKDIR /home/workspace

ADD . .

ENV PROD 1

RUN go build -buildmode=plugin -o _gameslib/fill.so _games/common.go _games/fill.go
RUN go build -o GoGame

EXPOSE 3001

CMD ./GoGame