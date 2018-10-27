FROM golang:1.10.3

WORKDIR $GOPATH/src/api-server
ADD . $GOPATH/src/api-server

RUN make

ENTRYPOINT ["./restful-api-server"]

EXPOSE 4000
