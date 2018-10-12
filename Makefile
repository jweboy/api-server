all: gotool
	go build -v .
clean:
	rm -f restful-api-server
gotool:
	gofmt -w .
help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"

.PHONY: clean gotool help