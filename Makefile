BINARY_NAME=eve-online-tools
DIST_FOLDER=dist

release: clean build

build: main.go
	GOARCH=amd64 GOOS=linux go build -o ./${DIST_FOLDER}/${BINARY_NAME} main.go 

clean:
	rm -Rf ./${DIST_FOLDER}