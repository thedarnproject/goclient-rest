BIN_DIR := bin
CLIENT_DIR := client

default: binary

binary:
	mkdir -p bin
	go build -o ${BIN_DIR}/goclient-rest ${CLIENT_DIR}/main.go

install:
	cd ${CLIENT_DIR}/ && go install

clean:
	rm -rf bin/

glide:
	glide update --strip-vendor

glide-vc:
	glide-vc --use-lock-file --only-code --no-tests

glide-hard:
	glide cache-clear
	glide update --strip-vendor

