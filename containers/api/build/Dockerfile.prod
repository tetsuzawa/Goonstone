FROM golang:1.13-alpine

RUN mkdir -p /go/src/app
ADD . /go/src/app

WORKDIR /go/src/app
RUN  go install ./cmd/server

CMD ["server"]
