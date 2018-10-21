FROM golang:1.10.3-alpine

WORKDIR $GOPATH/src/api-server
ADD . $GOPATH/src/api-server

RUN make

ENTRYPOINT ["./restful-api-server"]

EXPOSE 4000

# FROM scratch

# WORKDIR $GOPATH/src/api-server
# ADD . $GOPATH/src/api-server
# RUN make

# EXPOSE 4000
# CMD ["./restful-api-server"]
