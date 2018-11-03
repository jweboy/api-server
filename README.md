# restful-api-server [![Build Status](https://travis-ci.org/jweboy/restfult-api-server.svg?branch=master)](https://travis-ci.org/jweboy/restfult-api-server)

## 开发
> go run main.go

## 后期开发命令行保存
> sudo /usr/local/bin/docker-compose up
> sudo docker build -t jweboy/api-server:latest .
> sudo docker run -p 4000:4000 -d --name api-server --restart=always api-server
> sudo docker run --link mysql:mysql -p 4000:4000 api-server
> govendor add +local
> govendor add +external
> sudo docker run -it --rm -d --name apiserver -p 4000:4000 jweboy/apiserver

## TODO

- git push之前写一个shell脚本，保证vendor依赖添加完整，push完成之后删除vendor


## 参考
- [qiniu-sdk源代码](https://github.com/qiniu/api.v7/blob/master/storage/form_upload.go)
- [qiniu-sdk例子](https://github.com/qiniu/api.v7/blob/master/examples/form_upload_simple.go)
- [qiniu-sdk文档](https://developer.qiniu.com/kodo/sdk/1289/nodejs#server-upload)
[gin](https://github.com/gin-gonic/gin)
[Golang应用部署到Docker](https://segmentfault.com/a/1190000013960558#articleHeader3)