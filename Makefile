.PHONY: rm,build
all:rm build
rm:
	- rm -rf bin/
build:
	- CGO_ENABLED=0 GOOS=linux go build -ldflags '-d -w -s ' -a -installsuffix cgo -o bin/app .
test:
	- go test
docker:
	- docker build -t storage-api .