FROM golang:1.19.4-buster

WORKDIR /go/src/

COPY . /go/src/

RUN apt-get update

RUN go clean --modcache
RUN go mod download
RUN go build -o bars-fun-api main.go

CMD ["./bars-fun-api"]