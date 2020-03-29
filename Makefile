.PHONY: clean build

all: clean build

clean:
	rm -rf ./bin/darwin64/parseSlackMessage
	rm -rf ./bin/windows64/parseSlackMessage.exe

build:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin64/parseSlackMessage ./parseSlackMessage.go
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows64/parseSlackMessage.exe ./parseSlackMessage.go