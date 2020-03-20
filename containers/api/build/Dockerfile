FROM golang:1.13-alpine

RUN mkdir -p /go/src/app
ADD . /go/src/app

WORKDIR /go/src/app
RUN apk add git --update-cache

RUN go get -u github.com/yyh-gl/realize

CMD ["realize", "start"]
