
BINARY_NAME=GoldWatcher.app
APP_NAME=GoldWatcher

VERSION=1.0.1
BUILD_NO=2

## build: build binary and package app
build:
	
	fyne package -appVersion ${VERSION} -appBuild ${BUILD_NO} -name ${APP_NAME} -release -appID  ${BINARY_NAME} 
	

## run: builds and runs the application
run:
	env DB_PATH="./sql.db" go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleaned!"

## test: runs all tests
test:
	go test -v ./...