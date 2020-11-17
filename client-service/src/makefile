# Goパラメータ
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SRCCODE_NAME=client-service.go
BINARY_NAME=client-service
BINARY_LINUX=$(BINARY_NAME)_linux

# 実行するターゲット - buildのみ
all: build-linux
build:
	# -oは場所, -vは実行ファイル名
	CGO_ENABLED=0 $(GOBUILD) -ldflags="-s -w" -trimpath -o ../exe/$(BINARY_NAME) ./$(SRCCODE_NAME)
	upx ../exe/$(BINARY_NAME)
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/gin-gonic/gin


# クロスコンパイル
build-linux:
	GOOS=linux CGO_ENABLED=0 GOARCH=386 $(GOBUILD) -ldflags="-s -w" -trimpath -o ../exe/$(BINARY_LINUX) ./$(SRCCODE_NAME)
	upx ../exe/$(BINARY_LINUX)
docker-build:
	docker build -t sample-mesh/api-service -f ./docker/Dockerfile .
