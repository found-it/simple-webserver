FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/simple-webserver
COPY ./src/ .
COPY ./go.mod .

RUN go install -v .

ENTRYPOINT ["/go/bin/simple-webserver"]
