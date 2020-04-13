FROM golang:latest

WORKDIR /go/src/simple-webserver

COPY ./src .
COPY ./go.mod .

RUN go install -v .

# EXPOSE 8081

ENTRYPOINT ["simple-webserver"]
