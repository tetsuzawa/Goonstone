FROM golang:latest

RUN mkdir -p /go/src/Goonstone
WORKDIR /go/src/Goonstone
ADD . /go/src/Goonstone

RUN go build -o app ./cmd
RUN go get github.com/oxequa/realize

#CMD ["./app"]
CMD ["go", "run", "cmd/main.go"]
