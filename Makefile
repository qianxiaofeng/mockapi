all: build build-linux

build:
	go build -o .build/mockapi mockapi
	chmod +x .build/mockapi


build-linux:
	env GOOS=linux GOARCH=amd64 go build -o .build/linux/mockapi mockapi
	chmod +x .build/linux/mockapi


release: build-linux
	scp -i ~/aws/key/eospro.pem .build/linux/mockapi ubuntu@13.251.3.82:/home/ubuntu/mockapi/mockapi


PHONY: build build-linux release