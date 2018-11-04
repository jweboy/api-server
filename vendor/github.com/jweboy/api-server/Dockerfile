FROM golang
WORKDIR $GOPATH/src/github.com/jweboy/api-server
ADD . $GOPATH/src/github.com/jweboy/api-server
RUN make

EXPOSE 4000
ENTRYPOINT ["./api-server"]
