FROM golang:latest

WORKDIR $GOPATH/src/hangmango-web-api
ADD . $GOPATH/src/hangmango-web-api
RUN go build .

EXPOSE 8080

ENTRYPOINT ["./hangmango-web-api"]
