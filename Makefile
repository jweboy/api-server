SHELL := /bin/bash
BASEDIR = $(shell pwd)

# build with version info
versionDir="github.com/jweboy/api-server/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%Z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags = "-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

all: gotool
	@echo "============= build local API ============="
	go build -v -ldflags ${ldflags} .
clean:
	@echo "============= clean API file ============="
	rm -f api-server && rm -rf vendor
gotool:
	@echo "============= verdor local API ============="
	gofmt -w .
	go tool vet . |& grep -v verdor;true
help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"

# Docker opera
docker-build-image:
	@echo "============= docker build local image ============="
	sudo docker build -t jweboy/apiserver:latest .
docker-push-image:
	@echo "============= docker push this image ============="
	sudo docker push jweboy/apiserver:latest
docker-run-container:
	@echo "============= docker run it in container ============="
	sudo docker run -p 4000:4000 -d --name apiserver --restart=always jweboy/apiserver
.PHONY: clean gotool