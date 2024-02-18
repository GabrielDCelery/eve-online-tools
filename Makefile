BINARY_NAME=eve-online-tools
BUILD_FOLDER=build
DIST_FOLDER=dist

release: build dist

build: main.go
	rm -Rf ./${BUILD_FOLDER}
	mkdir ${BUILD_FOLDER}
	GOARCH=amd64 GOOS=linux go build -o ./${BUILD_FOLDER}/${BINARY_NAME}-linux main.go

dist:
ifndef RELEASE_TAG
$(error RELEASE_TAG is not set)
endif
	rm -Rf ./${DIST_FOLDER}
	mkdir ${DIST_FOLDER}
	cp ./${BUILD_FOLDER}/${BINARY_NAME}-linux ./${DIST_FOLDER}/${BINARY_NAME}
	cd ${DIST_FOLDER} && zip ./${BINARY_NAME}-${RELEASE_TAG}-linux.zip ./${BINARY_NAME}
	rm ./${DIST_FOLDER}/${BINARY_NAME}
