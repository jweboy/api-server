FROM golang

WORKDIR $GOPATH/src/restful-api-server
COPY . $GOPATH/src/restful-api-server
RUN make

ENTRYPOINT ["./restful-api-server"]

EXPOSE 4000