# restful-api-server

# 开发
> go run main.go


> sudo /usr/local/bin/docker-compose up
> sudo docker build -t jweboy/api-server:latest .
> sudo docker run --link mysql:mysql -p 4000:4000 -d --name api-server --restart=always api-server
> sudo docker run --link mysql:mysql -p 4000:4000 api-server

[https://segmentfault.com/a/1190000013960558#articleHeader3](https://segmentfault.com/a/1190000013960558#articleHeader3)