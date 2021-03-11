.PHONY: build
APP_NAME=confluence-cli

go-init:
	go mod init github.com/fr123k/confluence-client
	go mod vendor

build:
	go build -o build/${APP_NAME} cmd/${APP_NAME}.go
	go test -v --cover ./...

run: build
	./build/${APP_NAME}

clean:
	rm -rfv ./build
