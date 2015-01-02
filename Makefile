all: install test

clean:
	rm md5check

build:
	go build

install:
	go install

test:
	go fmt
	go vet
