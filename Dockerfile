FROM golang

WORKDIR $GOPATH/src/github.com/jweboy/restful-api-server
COPY . $GOPATH/src/github.com/jweboy/restful-api-server
RUN make

ENTRYPOINT ["./restful-api-server"]

EXPOSE 4000