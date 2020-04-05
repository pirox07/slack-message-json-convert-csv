.PHONY: clean build

all: clean build

clean:
	rm -rf ./bin/darwin64/slack-message-json-convert-csv
	rm -rf ./bin/windows64/slack-message-json-convert-csv.exe

build:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin64/slack-message-json-convert-csv ./main.go
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows64/slack-message-json-convert-csv.exe ./main.go