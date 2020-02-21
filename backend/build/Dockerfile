FROM golang:1.13-buster

RUN mkdir -p /go/src/Goonstone
ADD . /go/src/Goonstone

WORKDIR /go/src/Goonstone
RUN go build -o ./server ./cmd/server

CMD ["./server"]
