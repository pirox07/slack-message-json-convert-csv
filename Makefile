.PHONY: clean build

all: clean build

clean:
	rm -rf ./tool/parseSlackMessage

build:
	GOOS=darwin GOARCH=amd64 go build -o ./tool/parseSlackMessage ./parseSlackMessage.go