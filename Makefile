BUILD_DIR = ./build
BINARY_NAME = mtt
ENTRYPOINT = ./main.go 
 
all: build test
 
build:
	go build -o ./${BUILD_DIR}/${BINARY_NAME} ${ENTRYPOINT}
 
test:
	go test -v -cover ./...
 
run:
	./${BUILD_DIR}/${BINARY_NAME}
 
clean:
	go clean
	rm -rf ${BUILD_DIR}