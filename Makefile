.PHONY: build
build: 
	GOOS=linux GOARCH=amd64 go build -o bin/main main.go

package: build
	zip deployment.zip bin/main

clean:
	rm -rf ./bin ./deployment.zip