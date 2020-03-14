.PHONY: clean build

all: clean build

clean:
	rm -rf ./parseSlackMessage

build:
	GOOS=darwin GOARCH=amd64 go build -o ./parseSlackMessage ./parseSlackMessage.go