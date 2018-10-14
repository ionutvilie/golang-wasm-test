BINARY = test
GOARCH = amd64

VERSION?=?
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GITURL=$(shell git config --get remote.origin.url | sed "s/git@//g;s/\.git//g;s/:/\//g")

CURRENT_DIR=$(shell pwd)
BUILD_DIR_LINK=$(shell readlink ${BUILD_DIR})

# DOCKER_IMAGE_NAME       ?= ${BINARY}
# DOCKER_IMAGE_TAG        ?= latest

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-w -s -X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# Build the project
all: linux docker

# clean: 
# 	go clean -n
# 	rm -f ${CURRENT_DIR}/${BINARY}-windows-${GOARCH}.exe
# 	rm -f ${CURRENT_DIR}/${BINARY}-linux-${GOARCH}

linux:
	@echo ">> building linux binary" ;
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-linux-${GOARCH} . ;

wasm:
	@echo ">> building web binary" ;
	cd wasm-src && GOOS=js GOARCH=wasm go build -o ../build/main.wasm ;

# windows:
# 	@echo ">> building windows binary"
# 	GOOS=windows GOARCH=amd64 go build -o ${BINARY}-windows-${GOARCH}.exe . ;

# docker-build: 
# 	@echo ">> building docker image"
# 	docker build -t "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" -f DockerFile .;
# 	@echo ">> docker run --rm -p 8081:8081 --name ${BINARY} -d ${BINARY}" ;

# docker-run: 
# 	@echo ">> run docker image" ;
# 	docker run --rm -p 8080:8080 --name ${BINARY} -d ${BINARY} ;

release: linux docker

.PHONY: release all linux windows docker