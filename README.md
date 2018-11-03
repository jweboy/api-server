# restful-api-server [![Build Status](https://travis-ci.org/jweboy/restfult-api-server.svg?branch=master)](https://travis-ci.org/jweboy/restfult-api-server)

# 开发
> go run main.go


> sudo /usr/local/bin/docker-compose up
> sudo docker build -t jweboy/api-server:latest .
> sudo docker run -p 4000:4000 -d --name api-server --restart=always jweboy/apiserver
> sudo docker run --link mysql:mysql -p 4000:4000 api-server
> govendor add +local
> govendor add +external
> sudo docker run -it --rm -d --name apiserver -p 4000:4000 jweboy/apiserver

[https://segmentfault.com/a/1190000013960558#articleHeader3](https://segmentfault.com/a/1190000013960558#articleHeader3)