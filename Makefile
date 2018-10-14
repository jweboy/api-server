SHELL := /bin/bash
BASEDIR = $(shell pwd)

# build with version info
versionDir="github.com/jweboy/restful-api-server/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%Z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags = "-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

all: gotool
	go build -v -ldflags ${ldflags} .
clean:
	rm -f github.com/jweboy/restful-api-server
gotool:
	gofmt -w .
	go tool vet . |& grep -v verdor;true
help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"

.PHONY: clean gotool help