PLATFORMS := linux_amd64 linux_386 linux_arm5 linux_arm6 linux_arm7 linux_arm64 darwin_amd64 darwin_386 freebsd_amd64 freebsd_386 windows_386 windows_amd64 linux_arm64

NAME = github-release-tool
RELEASE_VERSION = 0.0.1

FLAGS_all = GOROOT=$(GOROOT) GOPATH=$(GOPATH)
FLAGS_linux_amd64   = $(FLAGS_all) 		GOOS=linux   GOARCH=amd64
FLAGS_linux_386     = $(FLAGS_all) 		GOOS=linux   GOARCH=386
FLAGS_linux_arm5     = $(FLAGS_all) 	GOOS=linux   GOARCH=arm   GOARM=5 # ARM5 support for Raspberry Pi
FLAGS_linux_arm6     = $(FLAGS_all) 	GOOS=linux   GOARCH=arm   GOARM=6 # ARM5 support for Raspberry Pi
FLAGS_linux_arm7     = $(FLAGS_all) 	GOOS=linux   GOARCH=arm   GOARM=7 # ARM5 support for Raspberry Pi
FLAGS_linux_arm7     = $(FLAGS_all) 	GOOS=linux   GOARCH=arm   GOARM=7 # ARM5 support for Raspberry Pi
FLAGS_linux_arm64     = $(FLAGS_all) 	GOOS=linux   GOARCH=arm64 		  # ARM5 support for Raspberry Pi
FLAGS_darwin_amd64  = $(FLAGS_all) 		GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0
FLAGS_darwin_386    = $(FLAGS_all) 		GOOS=darwin  GOARCH=386   CGO_ENABLED=0
FLAGS_freebsd_amd64  = $(FLAGS_all) 	GOOS=freebsd GOARCH=amd64 CGO_ENABLED=0
FLAGS_freebsd_386    = $(FLAGS_all) 	GOOS=freebsd GOARCH=386   CGO_ENABLED=0
FLAGS_windows_386   = $(FLAGS_all) 		GOOS=windows GOARCH=386   CGO_ENABLED=0
FLAGS_windows_amd64 = $(FLAGS_all) 		GOOS=windows GOARCH=amd64 CGO_ENABLED=0

.PHONY: clean test lint fmt build


deps:
	go get github.com/Unknwon/bra
	go get -d ./...

lint:
	go get github.com/golang/lint
	golint '*.go' '**/*.go'

run:
	./bin/github-release-tool

fmt:
	gofmt -w -s .

test:
	godep go test ./...

clean:
	rm -rf bin

build-all: clean $(foreach PLATFORM,$(PLATFORMS),build-$(PLATFORM))

build-%:
	echo "Compiling release for $*"
	$(FLAGS_$*) go build -ldflags "-X main.version=${RELEASE_VERSION}" -o 'bin/${NAME}-v${RELEASE_VERSION}-$*'

build:
	go build -o 'bin/${NAME}'
