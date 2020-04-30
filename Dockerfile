FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/simple-server
COPY ./src/ .
COPY ./go.mod .

RUN go install -v .

# RUN addgroup --gid 2323 "foundit" && \
#     adduser --disabled-password \
#             --home "/home/foundit" \
#             --ingroup "foundit" \
#             --no-create-home \
#             --uid 2324 \
#             "foundit"
#
# USER foundit

ENTRYPOINT ["/go/bin/simple-webserver"]
