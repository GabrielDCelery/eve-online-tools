BINARY_NAME=eve-online-tools
DIST_FOLDER=dist
ASSETS_FOLDER=assets

build-release-assets: build zip

build: main.go
	rm -Rf ./${DIST_FOLDER}
	mkdir ${DIST_FOLDER}
	GOARCH=amd64 GOOS=linux go build -o ./${DIST_FOLDER}/${BINARY_NAME}-linux main.go

zip:
ifndef RELEASE_TAG
$(error RELEASE_TAG is not set)
endif
	rm -Rf ./${ASSETS_FOLDER}
	mkdir ${ASSETS_FOLDER}
	cp ./${DIST_FOLDER}/${BINARY_NAME}-linux ./${ASSETS_FOLDER}/${BINARY_NAME}
	zip ./${ASSETS_FOLDER}/${BINARY_NAME}-${RELEASE_TAG}-linux.zip ./${ASSETS_FOLDER}/${BINARY_NAME}
	rm ./${ASSETS_FOLDER}/${BINARY_NAME}