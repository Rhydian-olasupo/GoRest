PROJECT_NAME = bookServer
BINARY_NAME = app
GOCMD=go
GOBUILD=$(GOCMD) build
build:
	$(info Building the bookServer Binary.....)
	cd ${PROJECT_NAME} && GOOS=linux GOARCH=arm ${GOBUILD} -o "${BINARY_NAME}" -v