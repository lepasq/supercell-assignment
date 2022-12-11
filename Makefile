BINARY_NAME=my_program

all: build test

build:
	./build.sh ${BINARY_NAME}
 
test:
	./test.sh
