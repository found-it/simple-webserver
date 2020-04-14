# FROM debian:7
#
# WORKDIR /go/src/simple-webserver
#
# COPY ./go1.14.2.linux-amd64.tar.gz /tmp
# COPY ./src .
# COPY ./go.mod .
#
# RUN cd /tmp && tar -C /usr/local -xzf go1.14.2.linux-amd64.tar.gz
# ENV GOBIN="/go/bin"
# ENV PATH="${PATH}:/usr/local/go/bin:${GOBIN}"
#
# RUN go install -v .
#
# # EXPOSE 8081
#
# ENTRYPOINT ["simple-webserver"]
# ENTRYPOINT ["sleep", "10000"]

FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/simple-webserver
COPY ./src/ .
COPY ./go.mod .

RUN go install -v .

FROM scratch

COPY --from=builder /go/bin/simple-webserver /go/bin/simple-webserver

ENTRYPOINT ["/go/bin/simple-webserver"]
