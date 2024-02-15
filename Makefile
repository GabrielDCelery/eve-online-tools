BINARY_NAME=eve-online-tools
DIST_FOLDER=dist
ASSETS_FOLDERS=assets

release: clean build zip

clean:
	rm -Rf ./${DIST_FOLDER}
	rm -Rf ./${ASSETS_FOLDERS}

build: main.go
	GOARCH=amd64 GOOS=linux go build -o ./${DIST_FOLDER}/${BINARY_NAME} main.go

zip:
	mkdir ${ASSETS_FOLDERS}
	cp ./${DIST_FOLDER}/* ./${ASSETS_FOLDERS}/
